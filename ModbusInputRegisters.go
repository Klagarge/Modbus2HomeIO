package main

import (
	"Modbus2HomeIO/homeio"
	"github.com/simonvetter/modbus"
)

func (m *ModbusHandler) HandleInputRegisters(req *modbus.InputRegistersRequest) (res []uint16, err error) {
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
		value := uint16(0)
		switch room {
		case homeio.LivingRoom:
			switch address {
			case 0, 1:
				value, err = m.home.GetShutterPosition(homeio.LivingRoom, homeio.Shutters1+homeio.Shutters(address-0))
			case 2:
				value, err = m.home.GetBrightness(homeio.LivingRoom)
			case 3:
				value, err = m.home.GetTemperature(homeio.LivingRoom)
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
