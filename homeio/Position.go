package homeio

import "fmt"

/* Home interface implementation */

func (h *home) GetPosition() (float64, float64, error) {
	// Read input values from inputs map and check if they exist.
	lat, err := h.inputs["lat"]
	if !err {
		return 0, 0, fmt.Errorf("latitude not found")
	}
	long, err := h.inputs["long"]
	if !err {
		return 0, 0, fmt.Errorf("longitude not found")
	}

	// Return input values.
	return lat.(float64), long.(float64), nil
}
