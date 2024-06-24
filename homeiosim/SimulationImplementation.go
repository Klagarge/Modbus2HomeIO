package homeiosim

import (
	"Modbus2HomeIO/homeio"
	"Modbus2HomeIO/nrgsim"
	"fmt"
	"math"
	"time"
)

type (
	simulation struct {
		home homeio.Home

		simulation nrgsim.Simulation
		pv         nrgsim.Measurable

		storage struct {
			controller StorageController
			device     nrgsim.Storable
		}
		heating  nrgsim.MeasurableGroup
		gates    nrgsim.MeasurableGroup
		shutters nrgsim.MeasurableGroup
		lights   nrgsim.MeasurableGroup
		devices  nrgsim.Measurable
		lastStep time.Time
	}
)

// New creates a new simulation with the given home. The simulation is initialized with all consumers and producers of the home, and the simulation is ready to be stepped forward.
func New(home homeio.Home) Simulation {
	s := &simulation{
		home:       home,
		simulation: nrgsim.New(),
		heating:    nrgsim.MeasurableGroup{},
		gates:      nrgsim.MeasurableGroup{},
		shutters:   nrgsim.MeasurableGroup{},
		lights:     nrgsim.MeasurableGroup{},
	}

	var err error
	err = home.Poll()
	if err != nil {
		panic(err)
	}
	s.lastStep, err = home.GetTime()
	if err != nil {
		panic(err)
	}

	/* Add PV production */

	s.pv, _ = s.simulation.AddAdjustableProducer(20000, func() float64 {
		brightness, err := home.GetBrightness(homeio.Exterior)
		if err != nil {
			return 0
		}
		return math.Max((float64(brightness)-5)*1.1/100.0, 0)
	})

	/* Add energy storage */

	s.storage.controller = StorageController{
		Mode:              nrgsim.StorageCharge,
		MaxChargePower:    5000,
		MaxDischargePower: 5000,
	}

	s.storage.device, _ = s.simulation.AddEnergyStorage(5000, 5000, 22*3.6e6, s.storage.controller.ControlFunction)

	/* Add heating consumers */

	for _, room := range []homeio.Room{homeio.LivingRoom, homeio.Kitchen, homeio.EntranceHall, homeio.BedroomCorridor, homeio.ChildrenRoom, homeio.Bathroom, homeio.SingleBedroom,
		homeio.PrivateBathroom, homeio.CoupleBedroom, homeio.LaundryRoom, homeio.HomeOffice} {
		s.heating[room.String()], _ = s.simulation.AddCustomConsumer(func() float64 {
			pwr, err := home.GetHeatingPowerOutput(room)
			if err != nil {
				return 0
			}
			return float64(pwr)
		})
	}

	/* Add entry gate and garage door */

	for _, room := range []homeio.Room{homeio.Exterior, homeio.Garage} {
		s.gates[room.String()], _ = s.simulation.AddControlledConsumer(1000, func() bool {
			dir, err := home.GetGateDirectionOutput(room)
			if err != nil {
				return false
			}
			return dir != homeio.GateStop
		})
	}

	/* Add shutters of living room (4) */

	for _, shutter := range []homeio.Shutters{homeio.Shutters1, homeio.Shutters2, homeio.Shutters3, homeio.Shutters4} {
		s.gates[homeio.LivingRoom.String()], _ = s.simulation.AddControlledConsumer(10, func() bool {
			dir, err := home.GetShuttersDirectionOutput(homeio.LivingRoom, shutter)
			if err != nil {
				return false
			}
			onTop, err := home.AreShuttersOnTop(homeio.LivingRoom, shutter)
			if err != nil {
				return false
			}
			onBottom, err := home.AreShuttersOnBottom(homeio.LivingRoom, shutter)
			if err != nil {
				return false
			}
			return dir == homeio.ShuttersUp && !onTop || dir == homeio.ShuttersDown && !onBottom
		})
	}

	/* Add shutters of other rooms */

	for _, room := range []homeio.Room{homeio.Kitchen, homeio.EntranceHall, homeio.Garage, homeio.BedroomCorridor, homeio.ChildrenRoom, homeio.SingleBedroom,
		homeio.CoupleBedroom, homeio.LaundryRoom, homeio.HomeOffice} {
		s.shutters[room.String()], _ = s.simulation.AddControlledConsumer(10, func() bool {
			dir, err := home.GetShuttersDirectionOutput(room, homeio.Shutters1)
			if err != nil {
				return false
			}
			onTop, err := home.AreShuttersOnTop(room, homeio.Shutters1)
			if err != nil {
				return false
			}
			onBottom, err := home.AreShuttersOnBottom(room, homeio.Shutters1)
			if err != nil {
				return false
			}
			return dir == homeio.ShuttersUp && !onTop || dir == homeio.ShuttersDown && !onBottom
		})
	}

	/* Add light consumers */

	/*s.lights[homeio.LivingRoom.String()], _ = s.simulation.AddAdjustableConsumer(60, func() float64 {
		level, err := home.GetDimmerPercentageOutput(homeio.LivingRoom, homeio.Light1)
		if err != nil {
			return 0
		}
		return float64(level) / 100.0
	})*/

	/* Add some noise consumers */

	//s.devices, _ = s.simulation.AddNoiseConsumer(25, 300)

	return s
}

/* Simulation interface implementation */

func (s *simulation) GetGridConnection() nrgsim.Measurable {
	return s.simulation.GetGridConnection()
}

func (s *simulation) GetPVProducer() nrgsim.Measurable {
	return s.pv
}

func (s *simulation) GetStorageController() *StorageController {
	return &s.storage.controller
}

func (s *simulation) GetStorage() nrgsim.Storable {
	return s.storage.device
}

func (s *simulation) GetHeatingConsumer(room homeio.Room) (nrgsim.Measurable, error) {
	if heating, ok := s.heating[room.String()]; ok {
		return heating, nil
	}
	return nil, fmt.Errorf("room %v not found", room)
}

func (s *simulation) GetHeatingConsumersTotal() nrgsim.Measurable {
	return s.heating
}

func (s *simulation) Step() {
	s.simulation.CalculateStep(time.Since(s.lastStep))
	t, err := s.home.GetTime()
	if err != nil {
		panic(err)
	}
	s.lastStep = t
}

func (s *simulation) Reset() {
	s.simulation.Reset()
	s.lastStep, _ = s.home.GetTime()
}
