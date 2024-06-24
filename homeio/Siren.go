package homeio

import (
	"fmt"
	"net/http"
)

/* Home interface implementation */

func (h *home) SetSirenOn(room Room, on bool) error {
	// Store output value in outputs map.
	h.outputs[fmt.Sprintf("siren/%s", room)] = on

	// Define siren name based on the room.
	siren := ""
	switch room {
	case EntranceHall:
		siren = "interior_sirene"
	case Exterior:
		siren = "exterior_sirene"
	default:
		return fmt.Errorf("invalid room")
	}

	// Send request to the server and check response status.
	response, err := h.client.Get(fmt.Sprintf("%s/stsw/%s/%s", h.baseURL, map[bool]string{true: "turn_on", false: "turn_off"}[on], siren))
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", response.Status)
	}

	// Return nil if everything is ok.
	return nil
}

func (h *home) GetSirenOnOutput(room Room) (bool, error) {
	// Read output value from outputs map and check if it exists.
	value, ok := h.outputs[fmt.Sprintf("siren/%s", room)]
	if !ok {
		return false, fmt.Errorf("siren not found")
	}

	// Return output value.
	return value.(bool), nil
}
