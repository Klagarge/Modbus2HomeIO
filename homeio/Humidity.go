package homeio

import "fmt"

func (h *home) GetRelativeHumidity() (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, err := h.inputs["rhm"]
	if !err {
		return 0, fmt.Errorf("relative humidity sensor not found")
	}

	// Return input value.
	return uint16(value.(float64) * 100), nil
}
