package main

import (
	"Modbus2HomeIO/homeio"
	"github.com/simonvetter/modbus"
)

func (m *ModbusHandler) HandleDiscreteInputs(req *modbus.DiscreteInputsRequest) (res []bool, err error) {
	room, err := homeio.UnitIDToRoom(req.UnitId)
	if err != nil {
		return
	}

	err = m.home.Poll()
	if err != nil {
		return
	}

	addresses := make([]uint16, req.Quantity)
	for i := uint16(0); i < req.Quantity; i++ {
		addresses[i] = req.Addr + i
	}

	for _, address := range addresses {
		value := false
		switch room {
		case homeio.LivingRoom:
			switch address {
			case 0, 1, 2:
				value, err = m.home.IsButtonPressed(homeio.LivingRoom, homeio.Button1+homeio.Button(address-0))
			case 3, 5, 7:
				value, err = m.home.IsDimmerControlUpPressed(homeio.LivingRoom, homeio.Dimmer1+homeio.Dimmer(address-3))
			case 4, 6, 8:
				value, err = m.home.IsDimmerControlDownPressed(homeio.LivingRoom, homeio.Dimmer1+homeio.Dimmer(address-4))
			case 9, 11:
				value, err = m.home.IsShutterControlUpPressed(homeio.LivingRoom, homeio.Shutters1+homeio.Shutters(address-9))
			case 10, 12:
				value, err = m.home.IsShutterControlDownPressed(homeio.LivingRoom, homeio.Shutters1+homeio.Shutters(address-10))
			case 13, 14:
				value, err = m.home.IsDoorOpen(homeio.LivingRoom, homeio.Door1+homeio.Door(address-13))
			case 15:
				value, err = m.home.IsSmokeDetected(homeio.LivingRoom)
			case 16:
				value, err = m.home.IsMotionDetected(homeio.LivingRoom)
			case 17:
				value, err = m.home.IsLightDetected(homeio.LivingRoom)
			default:
			}
		}

		if err != nil {
			return nil, err
		}
		res = append(res, value)
	}

	return
}
