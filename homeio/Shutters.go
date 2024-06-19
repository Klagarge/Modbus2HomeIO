package homeio

import (
	"fmt"
	"net/http"
)

type (
	// Shutters identifies a shutter in the house.
	Shutters int

	ShuttersDirection string
)

const (
	Shutters1 Shutters = 1
	Shutters2 Shutters = 2
	Shutters3 Shutters = 3
	Shutters4 Shutters = 4

	ShuttersUp   ShuttersDirection = "up"
	ShuttersDown ShuttersDirection = "down"
	ShuttersStop ShuttersDirection = "stopped"
)

func (h *home) IsShutterControlUpPressed(room Room, shutter Shutters) (bool, error) {
	value, ok := h.values[fmt.Sprintf("udsw/%d/%s/up", shutter, room)]
	if !ok {
		return false, fmt.Errorf("switch not found")
	}
	return value.(bool), nil
}

func (h *home) IsShutterControlDownPressed(room Room, shutter Shutters) (bool, error) {
	value, ok := h.values[fmt.Sprintf("udsw/%d/%s/down", shutter, room)]
	if !ok {
		return false, fmt.Errorf("switch not found")
	}
	return value.(bool), nil
}

func (h *home) GetShutterPosition(room Room, shutter Shutters) (uint16, error) {
	value, er := h.values[fmt.Sprintf("rso/%d/%s", shutter, room)]
	if !er {
		return 0, fmt.Errorf("shutter not found")
	}
	return uint16(value.(float64) / 10.0 * 65535), nil
}

func (h *home) SetShuttersDirection(room Room, shutter Shutters, direction ShuttersDirection) error {
	response, err := h.client.Get(fmt.Sprintf("%s/strs/%s/%d/%s", h.baseURL, direction, shutter, room))
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", response.Status)
	}
	return nil
}
