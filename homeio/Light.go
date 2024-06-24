package homeio

import (
	"fmt"
	"net/http"
)

// Light identifies a light groups in a room. Up to 3 lights are supported in a room and up to 5 lights are supported in the exterior.
type Light int

const (
	// Light1 is the first light in a room.
	Light1 Light = 1

	// Light2 is the second light in a room.
	Light2 Light = 2

	// Light3 is the third light in a room.
	Light3 Light = 3

	// EntranceLight is the light located at the entrance.
	EntranceLight Light = 1

	// GardenLight are the light pillars in the garden.
	GardenLight Light = 2

	// PoolLight are the lights inside the pool.
	PoolLight Light = 3

	// ExternalLight1 is the first light located on the west side of the house.
	ExternalLight1 Light = 4

	// ExternalLight2 is the second light located on the south side of the house.
	ExternalLight2 Light = 5
)

/* Home interface implementation */

func (h *home) IsLightDetected(room Room) (bool, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("bdtc/%s", room)]
	if !ok {
		return false, fmt.Errorf("light sensor not found")
	}

	// Return input value.
	return value.(bool), nil
}

func (h *home) GetBrightness(room Room) (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("bgs/%s", room)]
	if !ok {
		return 0, fmt.Errorf("brightness sensor not found")
	}

	// Return input value.
	return uint16(value.(float64) / 10.0 * 100.0), nil
}

func (h *home) SetLightOn(room Room, light Light, on bool) error {
	// Set local state in outputs map.
	if on {
		h.outputs[fmt.Sprintf("stlg/%d/%s", light, room)] = 100
	} else {
		h.outputs[fmt.Sprintf("stlg/%d/%s", light, room)] = 0
	}

	// Construct the url depending on the room and light.
	url := ""
	switch room {
	case Exterior:
		switch light {
		case EntranceLight:
			url = fmt.Sprintf("%s/stsw/%s/entrance_lights", h.baseURL, map[bool]string{true: "turn_on", false: "turn_off"}[on])

		case GardenLight:
			url = fmt.Sprintf("%s/stsw/%s/garden_lights", h.baseURL, map[bool]string{true: "turn_on", false: "turn_off"}[on])

		case PoolLight:
			url = fmt.Sprintf("%s/stsw/%s/pool_lights", h.baseURL, map[bool]string{true: "turn_on", false: "turn_off"}[on])

		case ExternalLight1:
			url = fmt.Sprintf("%s/stsw/%s/porche_1_lights", h.baseURL, map[bool]string{true: "turn_on", false: "turn_off"}[on])

		case ExternalLight2:
			url = fmt.Sprintf("%s/stsw/%s/porche_2_lights", h.baseURL, map[bool]string{true: "turn_on", false: "turn_off"}[on])

		default:
			return fmt.Errorf("invalid light light")
		}
	default:
		url = fmt.Sprintf("%s/swl/%s/%d/%s", h.baseURL, map[bool]string{true: "turn_on", false: "turn_off"}[on], light, room)
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
