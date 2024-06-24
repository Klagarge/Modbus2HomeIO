package homeio

import (
	"fmt"
	"net/http"
)

// Dimmer identifies a dimmer in a room. Up to 3 dimmers are supported in a room
type Dimmer int

const (
	// Dimmer1 is the first dimmer in a room.
	Dimmer1 Dimmer = 1

	// Dimmer2 is the second dimmer in a room.
	Dimmer2 Dimmer = 2

	// Dimmer3 is the third dimmer in a room.
	Dimmer3 Dimmer = 3
)

/* Home interface implementation */

func (h *home) IsDimmerControlUpPressed(room Room, dimmer Dimmer) (bool, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("lsd/%d/%s/up", dimmer, room)]
	if !ok {
		return false, fmt.Errorf("dimmer not found")
	}

	// Return input value.
	return value.(bool), nil
}

func (h *home) IsDimmerControlDownPressed(room Room, dimmer Dimmer) (bool, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("lsd/%d/%s/down", dimmer, room)]
	if !ok {
		return false, fmt.Errorf("dimmer not found")
	}

	// Return input value.
	return value.(bool), nil
}

func (h *home) SetDimmerPercentage(room Room, light Light, dim uint16) error {
	// Set local state in outputs map.
	h.outputs[fmt.Sprintf("stlg/%d/%s", light, room)] = dim

	// Construct the url depending on the room and light.
	url := ""
	switch room {
	case Exterior:
		switch light {
		case EntranceLight:
			url = fmt.Sprintf("%s/stlg/entrance_lights/%f", h.baseURL, float64(dim)/100.0*10.0)

		case GardenLight:
			url = fmt.Sprintf("%s/stlg/garden_lights/%f", h.baseURL, float64(dim)/100.0*10.0)

		case PoolLight:
			url = fmt.Sprintf("%s/stlg/pool_lights/%f", h.baseURL, float64(dim)/100.0*10.0)

		case ExternalLight1:
			url = fmt.Sprintf("%s/stlg/porche_1_lights/%f", h.baseURL, float64(dim)/100.0*10.0)

		case ExternalLight2:
			url = fmt.Sprintf("%s/stlg/porche_2_lights/%f", h.baseURL, float64(dim)/100.0*10.0)

		default:
			return fmt.Errorf("invalid light light")
		}
	default:
		url = fmt.Sprintf("%s/stl/%d/%s/%f", h.baseURL, light, room, float64(dim)/100.0*10.0)
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

func (h *home) GetDimmerPercentageOutput(room Room, light Light) (uint16, error) {
	// Get output value from outputs map and check if it exists.
	value, ok := h.outputs[fmt.Sprintf("stlg/%d/%s", light, room)]
	if !ok {
		return 0, fmt.Errorf("dimmer not found")
	}

	// Return output value.
	return value.(uint16), nil
}
