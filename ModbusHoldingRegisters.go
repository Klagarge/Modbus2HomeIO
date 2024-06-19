package main

import (
	"Modbus2HomeIO/homeio"
	"fmt"
	"github.com/simonvetter/modbus"
)

func (m *ModbusHandler) HandleHoldingRegisters(req *modbus.HoldingRegistersRequest) (res []uint16, err error) {
	room, err := homeio.UnitIDToRoom(req.UnitId)
	if err != nil {
		return
	}

	addresses := make([]uint16, req.Quantity)
	for i := uint16(0); i < req.Quantity; i++ {
		addresses[i] = req.Addr + i
	}

	if req.IsWrite {
		for _, address := range addresses {
			switch room {
			case homeio.LivingRoom:
				switch address {
				case 0:
					err = m.home.SetLightDim(homeio.LivingRoom, homeio.Light1, req.Args[address-req.Addr])
				}
			}
		}
	} else {
		return nil, fmt.Errorf("read access to holding registers is not supported")
	}
}
