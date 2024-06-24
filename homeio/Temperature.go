package homeio

import "fmt"

/* Home interface implementation */

func (h *home) GetTemperature(room Room) (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("temp/%s", room)]
	if !ok {
		return 0, fmt.Errorf("temperature sensor not found")
	}

	// Return input value.
	return uint16(value.(float64) * 1000), nil
}

func (h *home) GetTemperateSetPoint(room Room) (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("tsp/%s", room)]
	if !ok {
		return 0, fmt.Errorf("temperature set point not found")
	}

	// Return input value.
	return uint16(value.(float64) * 1000), nil
}

func (h *home) GetOutsideTemperature() (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs["otemp"]
	if !ok {
		return 0, fmt.Errorf("outside temperature sensor not found")
	}

	// Return input value.
	return uint16((value.(float64) - 273.15) * 1000), nil
}
