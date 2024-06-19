package homeio

import "fmt"

func (h *home) GetTemperature(room Room) (uint16, error) {
	value, ok := h.values[fmt.Sprintf("temp/%s", room)]
	if !ok {
		return 0, fmt.Errorf("temperature sensor not found")
	}
	return uint16(value.(float64) * 1000), nil
}

func (h *home) GetTemperateSetPoint(room Room) (float64, error) {
	value, ok := h.values[fmt.Sprintf("tsp/%s", room)]
	if !ok {
		return 0, fmt.Errorf("temperature set point not found")
	}
	return value.(float64), nil
}

func (h *home) GetOutsideTemperature() (float64, error) {
	value, err := h.values["otemp"]
	if !err {
		return 0, fmt.Errorf("outside temperature sensor not found")
	}
	return value.(float64), nil
}
