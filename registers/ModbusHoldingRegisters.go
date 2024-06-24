package registers

import (
	"Modbus2HomeIO/homeio"
	"Modbus2HomeIO/nrgsim"
	"fmt"
	"github.com/simonvetter/modbus"
)

func (m *Handler) HandleHoldingRegisters(req *modbus.HoldingRegistersRequest) (res []uint16, err error) {
	// UnitID 17 is the energy simulation, so handle it separately.
	if req.UnitId == 17 && req.IsWrite {
		for address := req.Addr; address < req.Addr+req.Quantity; address += 1 {
			switch address {
			case 0:
				switch req.Args[0] {
				case 65535:
					m.sim.GetStorageController().Mode = nrgsim.StorageDischarge
				case 1:
					m.sim.GetStorageController().Mode = nrgsim.StorageCharge
				default:
					m.sim.GetStorageController().Mode = nrgsim.StorageIdle
				}
			case 1:
				m.sim.GetStorageController().MaxChargePower = float64(req.Args[1])
			case 2:
				m.sim.GetStorageController().MaxDischargePower = float64(req.Args[2])
			}
		}
		return
	}

	// We only support writing to holding registers.
	if !req.IsWrite {
		return nil, fmt.Errorf("read access to holding registers is not supported")
	}

	// Get room from unit ID and fail if the room does not exist.
	room, err := homeio.UnitIDToRoom(req.UnitId)
	if err != nil {
		return []uint16{}, fmt.Errorf("invalid unit ID %v", req.UnitId)
	}

	for address := req.Addr; address < req.Addr+req.Quantity; address++ {
		switch room {
		case homeio.LivingRoom:
			switch address {
			case 0:
				err = m.home.SetDimmerPercentage(homeio.LivingRoom, homeio.Light1, req.Args[address-req.Addr])

			case 1:
				err = m.home.SetHeatingPower(homeio.LivingRoom, req.Args[address-req.Addr])
			}

		case homeio.GuestRestRoom:
			switch address {
			case 0, 1:
				err = m.home.SetDimmerPercentage(homeio.GuestRestRoom, homeio.Light1+homeio.Light(address-0), req.Args[address-req.Addr])
			}

		case homeio.Pantry:
			switch address {
			case 0:
				err = m.home.SetDimmerPercentage(homeio.Pantry, homeio.Light1, req.Args[address-req.Addr])
			}

		case homeio.Kitchen:
			switch address {
			case 0, 1:
				err = m.home.SetDimmerPercentage(homeio.Kitchen, homeio.Light1+homeio.Light(address-0), req.Args[address-req.Addr])
			case 2:
				err = m.home.SetHeatingPower(homeio.Kitchen, req.Args[address-req.Addr])
			}

		case homeio.EntranceHall:
			switch address {
			case 0:
				err = m.home.SetDimmerPercentage(homeio.EntranceHall, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetHeatingPower(homeio.EntranceHall, req.Args[address-req.Addr])
			}

		case homeio.Garage:
			switch address {
			case 0, 1:
				err = m.home.SetDimmerPercentage(homeio.Garage, homeio.Light1+homeio.Light(address-0), req.Args[address-req.Addr])
			}

		case homeio.BedroomCorridor:
			switch address {
			case 0:
				err = m.home.SetDimmerPercentage(homeio.BedroomCorridor, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetHeatingPower(homeio.BedroomCorridor, req.Args[address-req.Addr])
			}

		case homeio.ChildrenRoom:
			switch address {
			case 0:
				err = m.home.SetDimmerPercentage(homeio.ChildrenRoom, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetHeatingPower(homeio.ChildrenRoom, req.Args[address-req.Addr])
			}

		case homeio.Bathroom:
			switch address {
			case 0, 1:
				err = m.home.SetDimmerPercentage(homeio.Bathroom, homeio.Light1+homeio.Light(address-0), req.Args[address-req.Addr])
			case 2:
				err = m.home.SetHeatingPower(homeio.Bathroom, req.Args[address-req.Addr])
			}

		case homeio.SingleBedroom:
			switch address {
			case 0:
				err = m.home.SetDimmerPercentage(homeio.SingleBedroom, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetHeatingPower(homeio.SingleBedroom, req.Args[address-req.Addr])
			}

		case homeio.PrivateBathroom:
			switch address {
			case 0:
				err = m.home.SetDimmerPercentage(homeio.PrivateBathroom, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetHeatingPower(homeio.PrivateBathroom, req.Args[address-req.Addr])
			}

		case homeio.CoupleBedroom:
			switch address {
			case 0:
				err = m.home.SetDimmerPercentage(homeio.CoupleBedroom, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetHeatingPower(homeio.CoupleBedroom, req.Args[address-req.Addr])
			}

		case homeio.LaundryRoom:
			switch address {
			case 0:
				err = m.home.SetDimmerPercentage(homeio.LaundryRoom, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetHeatingPower(homeio.LaundryRoom, req.Args[address-req.Addr])
			}

		case homeio.HomeOffice:
			switch address {
			case 0, 1, 2:
				err = m.home.SetDimmerPercentage(homeio.HomeOffice, homeio.Light1+homeio.Light(address-0), req.Args[address-req.Addr])
			case 3:
				err = m.home.SetHeatingPower(homeio.HomeOffice, req.Args[address-req.Addr])
			}

		case homeio.Exterior:
			switch address {
			case 0, 1, 2, 3, 4:
				err = m.home.SetDimmerPercentage(homeio.Exterior, homeio.EntranceLight+homeio.Light(address-0), req.Args[address-req.Addr])
			}
		}

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
