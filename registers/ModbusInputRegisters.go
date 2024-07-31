package registers

import (
	"Modbus2HomeIO/homeio"
	"fmt"

	modbus "github.com/Klagarge/modbusGo"
)

func (m *Handler) HandleInputRegisters(req *modbus.InputRegistersRequest) (res []uint16, err error) {
	// UnitID 17 is the energy simulation, so handle it separately.
	if req.UnitId == 17 {

		// All values are float32, so quantity must be even.
		if req.Quantity%2 != 0 {
			return nil, fmt.Errorf("quantity must be even")
		}

		for address := req.Addr; address < req.Addr+req.Quantity; address += 2 {
			switch address {
			case 0:
				appendFloat32(&res, float32(m.sim.GetGridConnection().GetPower()))
			case 2:
				appendFloat32(&res, float32(m.sim.GetGridConnection().GetEnergy()))
			case 4:
				appendFloat32(&res, float32(m.sim.GetPVProducer().GetPower()))
			case 6:
				appendFloat32(&res, float32(m.sim.GetPVProducer().GetEnergy()))
			case 8:
				appendFloat32(&res, float32(m.sim.GetStorage().GetStateOfCharge()))
			case 10:
				appendFloat32(&res, float32(m.sim.GetStorage().GetPower()))
			case 12:
				appendFloat32(&res, float32(m.sim.GetHeatingConsumersTotal().GetPower()))
			case 14:
				appendFloat32(&res, float32(m.sim.GetHeatingConsumersTotal().GetEnergy()))
			}
		}

		return
	}

	// Get room from unit ID and fail if the room does not exist.
	room, err := homeio.UnitIDToRoom(req.UnitId)
	if err != nil {
		return
	}

	for address := req.Addr; address < req.Addr+req.Quantity; address++ {
		value := uint16(0)
		switch room {
		case homeio.LivingRoom:
			switch address {
			case 0, 1, 2, 3:
				value, err = m.home.GetShutterPosition(homeio.LivingRoom, homeio.Shutters1+homeio.Shutters(address-0))
			case 4:
				value, err = m.home.GetBrightness(homeio.LivingRoom)
			case 5:
				value, err = m.home.GetTemperature(homeio.LivingRoom)
			case 6:
				value, err = m.home.GetTemperateSetPoint(homeio.LivingRoom)
			default:
			}

		case homeio.Kitchen:
			switch address {
			case 0:
				value, err = m.home.GetShutterPosition(homeio.Kitchen, homeio.Shutters1)
			case 1:
				value, err = m.home.GetBrightness(homeio.Kitchen)
			case 2:
				value, err = m.home.GetTemperature(homeio.Kitchen)
			case 3:
				value, err = m.home.GetTemperateSetPoint(homeio.Kitchen)
			}

		case homeio.EntranceHall:
			switch address {
			case 0:
				value, err = m.home.GetShutterPosition(homeio.EntranceHall, homeio.Shutters1)
			case 1:
				value, err = m.home.GetBrightness(homeio.EntranceHall)
			case 2:
				value, err = m.home.GetTemperature(homeio.EntranceHall)
			case 3:
				value, err = m.home.GetTemperateSetPoint(homeio.EntranceHall)
			}

		case homeio.Garage:
			switch address {
			case 0:
				value, err = m.home.GetShutterPosition(homeio.Garage, homeio.Shutters1)
			case 1:
				value, err = m.home.GetBrightness(homeio.Garage)
			}

		case homeio.BedroomCorridor:
			switch address {
			case 0:
				value, err = m.home.GetShutterPosition(homeio.BedroomCorridor, homeio.Shutters1)
			case 1:
				value, err = m.home.GetTemperature(homeio.BedroomCorridor)
			case 2:
				value, err = m.home.GetTemperateSetPoint(homeio.BedroomCorridor)
			}

		case homeio.ChildrenRoom:
			switch address {
			case 0:
				value, err = m.home.GetShutterPosition(homeio.ChildrenRoom, homeio.Shutters1)
			case 1:
				value, err = m.home.GetBrightness(homeio.ChildrenRoom)
			case 2:
				value, err = m.home.GetTemperature(homeio.ChildrenRoom)
			case 3:
				value, err = m.home.GetTemperateSetPoint(homeio.ChildrenRoom)
			}

		case homeio.Bathroom:
			switch address {
			case 0:
				value, err = m.home.GetTemperature(homeio.Bathroom)
			case 1:
				value, err = m.home.GetTemperateSetPoint(homeio.Bathroom)
			}

		case homeio.SingleBedroom:
			switch address {
			case 0:
				value, err = m.home.GetShutterPosition(homeio.SingleBedroom, homeio.Shutters1)
			case 1:
				value, err = m.home.GetBrightness(homeio.SingleBedroom)
			case 2:
				value, err = m.home.GetTemperature(homeio.SingleBedroom)
			case 3:
				value, err = m.home.GetTemperateSetPoint(homeio.SingleBedroom)
			}

		case homeio.PrivateBathroom:
			switch address {
			case 0:
				value, err = m.home.GetTemperature(homeio.PrivateBathroom)
			case 1:
				value, err = m.home.GetTemperateSetPoint(homeio.PrivateBathroom)
			}

		case homeio.CoupleBedroom:
			switch address {
			case 0:
				value, err = m.home.GetShutterPosition(homeio.CoupleBedroom, homeio.Shutters1)
			case 1:
				value, err = m.home.GetBrightness(homeio.CoupleBedroom)
			case 2:
				value, err = m.home.GetTemperature(homeio.CoupleBedroom)
			case 3:
				value, err = m.home.GetTemperateSetPoint(homeio.CoupleBedroom)
			}

		case homeio.LaundryRoom:
			switch address {
			case 0:
				value, err = m.home.GetShutterPosition(homeio.LaundryRoom, homeio.Shutters1)
			case 1:
				value, err = m.home.GetBrightness(homeio.LaundryRoom)
			case 2:
				value, err = m.home.GetTemperature(homeio.LaundryRoom)
			case 3:
				value, err = m.home.GetTemperateSetPoint(homeio.LaundryRoom)
			}

		case homeio.HomeOffice:
			switch address {
			case 0:
				value, err = m.home.GetShutterPosition(homeio.HomeOffice, homeio.Shutters1)
			case 1:
				value, err = m.home.GetBrightness(homeio.HomeOffice)
			case 2:
				value, err = m.home.GetTemperature(homeio.HomeOffice)
			case 3:
				value, err = m.home.GetTemperateSetPoint(homeio.HomeOffice)
			}

		case homeio.Exterior:
			switch address {
			case 0:
				value, err = m.home.GetBrightness(homeio.Exterior)
			case 1:
				value, err = m.home.GetOutsideTemperature()
			case 2:
				value, err = m.home.GetRelativeHumidity()
			case 3:
				value, err = m.home.GetWindSpeed()
			}

		case homeio.Miscellaneous:
			switch address {
			case 0:
				value, err = m.home.GetYear()
			case 1:
				value, err = m.home.GetMonth()
			case 2:
				value, err = m.home.GetDay()
			case 3:
				value, err = m.home.GetHour()
			case 4:
				value, err = m.home.GetMinute()
			case 5:
				value, err = m.home.GetSecond()
			}
		}

		if err != nil {
			return nil, err
		}
		res = append(res, value)
	}

	return
}
