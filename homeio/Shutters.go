package homeio

import (
	"fmt"
	"net/http"
)

type (
	// Shutters identifies a shutter in a room. Up to 4 shutters are supported in a room.
	Shutters int

	// ShuttersDirection identifies the direction of the shutters. The shutters can be moved up, down, or stopped.
	ShuttersDirection string
)

const (
	// Shutters1 is the first shutter in a room.
	Shutters1 Shutters = 1

	// Shutters2 is the second shutter in a room.
	Shutters2 Shutters = 2

	// Shutters3 is the third shutter in a room.
	Shutters3 Shutters = 3

	// Shutters4 is the fourth shutter in a room.
	Shutters4 Shutters = 4

	// ShuttersUp is the direction to move the shutters up.
	ShuttersUp ShuttersDirection = "up"

	// ShuttersDown is the direction to move the shutters down.
	ShuttersDown ShuttersDirection = "down"

	// ShuttersStop is the direction to stop the shutters.
	ShuttersStop ShuttersDirection = "stopped"
)

/* Home interface implementation */

func (h *home) IsShutterControlUpPressed(room Room, shutter Shutters) (bool, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("udsw/%d/%s/up", shutter, room)]
	if !ok {
		return false, fmt.Errorf("switch not found")
	}

	// Return input value.
	return value.(bool), nil
}

func (h *home) IsShutterControlDownPressed(room Room, shutter Shutters) (bool, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("udsw/%d/%s/down", shutter, room)]
	if !ok {
		return false, fmt.Errorf("switch not found")
	}

	// Return input value.
	return value.(bool), nil
}

func (h *home) AreShuttersOnTop(room Room, shutter Shutters) (bool, error) {
	// Get the position of the shutters.
	p, err := h.GetShutterPosition(room, shutter)
	if err != nil {
		return false, err
	}

	// Check if the shutters are on top.
	return p == 65535, nil
}

func (h *home) AreShuttersOnBottom(room Room, shutter Shutters) (bool, error) {
	// Get the position of the shutters.
	p, err := h.GetShutterPosition(room, shutter)
	if err != nil {
		return false, err
	}

	// Check if the shutters are on bottom.
	return p == 0, nil
}

func (h *home) GetShutterPosition(room Room, shutter Shutters) (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, er := h.inputs[fmt.Sprintf("rso/%d/%s", shutter, room)]
	if !er {
		return 0, fmt.Errorf("shutters not found")
	}

	// Return input value.
	return uint16(value.(float64) / 10.0 * 1000), nil
}

func (h *home) SetShuttersDirection(room Room, shutter Shutters, direction ShuttersDirection) error {
	// Set local state in outputs map.
	h.outputs[fmt.Sprintf("strs/%d/%s", shutter, room)] = string(direction)

	// Send request and check response status.
	response, err := h.client.Get(fmt.Sprintf("%s/strs/%s/%d/%s", h.baseURL, direction, shutter, room))
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", response.Status)
	}

	// Return nil if everything is ok.
	return nil
}

func (h *home) GetShuttersDirectionOutput(room Room, shutter Shutters) (ShuttersDirection, error) {
	// Read output value from outputs map and check if it exists.
	value, ok := h.outputs[fmt.Sprintf("strs/%d/%s", shutter, room)]
	if !ok {
		return "", fmt.Errorf("shutters not found")
	}

	// Return output value.
	return ShuttersDirection(value.(string)), nil
}
