package homeio

import (
	"fmt"
	"net/http"
)

// Light identifies a light groups in the house.
type Light int

const (
	Light1         Light = 1
	Light2         Light = 2
	Light3         Light = 3
	EntranceLight  Light = 1
	GardenLight    Light = 2
	PoolLight      Light = 3
	ExternalLight1 Light = 4
	ExternalLight2 Light = 5
)

func (h *home) IsLightDetected(room Room) (bool, error) {
	value, ok := h.values[fmt.Sprintf("bdtc/%s", room)]
	if !ok {
		return false, fmt.Errorf("light sensor not found")
	}
	return value.(bool), nil
}

func (h *home) GetBrightness(room Room) (uint16, error) {
	value, ok := h.values[fmt.Sprintf("bgs/%s", room)]
	if !ok {
		return 0, fmt.Errorf("brightness sensor not found")
	}
	return uint16(value.(float64) / 10.0 * 65535), nil
}

func (h *home) SetLightOn(room Room, light Light, on bool) error {
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

	response, err := h.client.Get(url)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", response.Status)
	}
	return nil
}

func (h *home) SetLightDim(room Room, light Light, dim uint16) error {
	url := ""
	switch room {
	case Exterior:
		switch light {
		case EntranceLight:
			url = fmt.Sprintf("%s/stlg/entrance_lights/%f", h.baseURL, float64(dim)/65535.0*10.0)

		case GardenLight:
			url = fmt.Sprintf("%s/stlg/garden_lights/%f", h.baseURL, float64(dim)/65535.0*10.0)

		case PoolLight:
			url = fmt.Sprintf("%s/stlg/pool_lights/%f", h.baseURL, float64(dim)/65535.0*10.0)

		case ExternalLight1:
			url = fmt.Sprintf("%s/stlg/porche_1_lights/%f", h.baseURL, float64(dim)/65535.0*10.0)

		case ExternalLight2:
			url = fmt.Sprintf("%s/stlg/porche_2_lights/%f", h.baseURL, float64(dim)/65535.0*10.0)

		default:
			return fmt.Errorf("invalid light light")
		}
	default:
		url = fmt.Sprintf("%s/stl/%d/%s/%f", h.baseURL, light, room, float64(dim)/65535.0*10.0)
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
