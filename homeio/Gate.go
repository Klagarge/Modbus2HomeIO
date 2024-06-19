package homeio

import (
	"fmt"
	"net/http"
)

type GateDirection string

const (
	GateOpen  GateDirection = "open"
	GateClose GateDirection = "close"
	GateStop  GateDirection = "stop"
)

func (h *home) IsGateOpen(room Room) (bool, error) {
	switch room {
	case Exterior:
		value, ok := h.values["gts/entrance_gate/open"]
		if !ok {
			return false, fmt.Errorf("gate not found")
		}
		return value.(bool), nil

	case Garage:
		value, ok := h.values["gts/garage_door/open"]
		if !ok {
			return false, fmt.Errorf("gate not found")
		}
		return value.(bool), nil

	default:
		return false, fmt.Errorf("invalid room")
	}
}

func (h *home) IsGateClosed(room Room) (bool, error) {
	switch room {
	case Exterior:
		value, ok := h.values["gts/entrance_gate/closed"]
		if !ok {
			return false, fmt.Errorf("gate not found")
		}
		return value.(bool), nil

	case Garage:
		value := h.values["gts/garage_door/closed"]
		if value == nil {
			return false, fmt.Errorf("gate not found")
		}
		return value.(bool), nil

	default:
		return false, fmt.Errorf("invalid room")
	}
}

func (h *home) SetGateDirection(room Room, direction GateDirection) error {
	url := fmt.Sprintf("%s/cgate/%s", h.baseURL, direction)
	switch room {
	case Exterior:
		url += "/entrance_gate"

	case Garage:
		url += "/garage_door"

	default:
		return fmt.Errorf("invalid room")
	}

	response, err := h.client.Get(url)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", response.Status)
	}
	return nil
}
