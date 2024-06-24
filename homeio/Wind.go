package homeio

import "fmt"

/* Home interface implementation */

func (h *home) GetWindSpeed() (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, err := h.inputs["wdsp"]
	if !err {
		return 0, fmt.Errorf("wind speed sensor not found")
	}

	// Return input value.
	return uint16(value.(float64) * 1000), nil
}
