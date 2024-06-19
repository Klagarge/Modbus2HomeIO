package main

import (
	"Modbus2HomeIO/homeio"
)

type ModbusHandler struct {
	home homeio.Home
}

func NewModbusHandler(home homeio.Home) *ModbusHandler {
	return &ModbusHandler{
		home: home,
	}
}
