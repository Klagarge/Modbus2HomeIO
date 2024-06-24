package registers

import (
	"Modbus2HomeIO/homeio"
	"Modbus2HomeIO/homeiosim"
)

type Handler struct {
	home homeio.Home
	sim  homeiosim.Simulation
}

// NewHandler creates a new handler with the given home and electrical simulation.
func NewHandler(home homeio.Home, sim homeiosim.Simulation) *Handler {
	return &Handler{
		home: home,
		sim:  sim,
	}
}
