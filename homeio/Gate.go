package homeio

import (
	"fmt"
	"net/http"
)

// GateDirection identifies the movement direction of a gate. A gate can move to the direction open, close or be still in the case of stop.
type GateDirection string

const (
	// GateOpen is the direction to open a gate.
	GateOpen GateDirection = "open"

	// GateClose is the direction to close a gate.
	GateClose GateDirection = "close"

	// GateStop is the direction to stop a gate movement.
	GateStop GateDirection = "stop"
)

/* Home interface implementation */

func (h *home) IsGateOpen(room Room) (bool, error) {
	switch room {
	case Exterior:
		// Read input value from inputs map and check if it exists.
		value, ok := h.inputs["gts/entrance_gate/open"]
		if !ok {
			return false, fmt.Errorf("gate not found")
		}

		// Return input value.
		return value.(bool), nil

	case Garage:
		// Read input value from inputs map and check if it exists.
		value, ok := h.inputs["gts/garage_door/open"]
		if !ok {
			return false, fmt.Errorf("gate not found")
		}

		// Return input value.
		return value.(bool), nil

	default:
		// Return error if the room is invalid.
		return false, fmt.Errorf("invalid room")
	}
}

func (h *home) IsGateClosed(room Room) (bool, error) {
	switch room {
	case Exterior:
		// Read input value from inputs map and check if it exists.
		value, ok := h.inputs["gts/entrance_gate/closed"]
		if !ok {
			return false, fmt.Errorf("gate not found")
		}

		// Return input value.
		return value.(bool), nil

	case Garage:
		// Read input value from inputs map and check if it exists.
		value := h.inputs["gts/garage_door/closed"]
		if value == nil {
			return false, fmt.Errorf("gate not found")
		}

		// Return input value.
		return value.(bool), nil

	default:
		// Return error if the room is invalid.
		return false, fmt.Errorf("invalid room")
	}
}

func (h *home) SetGateDirection(room Room, direction GateDirection) error {
	// Set local state in outputs map.
	h.outputs[fmt.Sprintf("cgate/%s", room)] = direction

	// Construct the url depending on the room and direction.
	url := fmt.Sprintf("%s/cgate/%s", h.baseURL, direction)
	switch room {
	case Exterior:
		url += "/entrance_gate"

	case Garage:
		url += "/garage_door"

	default:
		return fmt.Errorf("invalid room")
	}

	// Send request and check response status.
	response, err := h.client.Get(url)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", response.Status)
	}

	// Return nil if everything went well.
	return nil
}

func (h *home) GetGateDirectionOutput(room Room) (GateDirection, error) {
	// Read output value from outputs map and check if it exists.
	value, ok := h.outputs[fmt.Sprintf("cgate/%s", room)]
	if !ok {
		return "", fmt.Errorf("gate not found")
	}

	// Return output value.
	return value.(GateDirection), nil
}
