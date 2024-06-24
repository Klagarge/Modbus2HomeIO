package homeio

import "fmt"

/* Home interface implementation */

func (h *home) IsMotionDetected(room Room) (bool, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("mdtc/%s", room)]
	if !ok {
		return false, fmt.Errorf("motion sensor not found")
	}

	// Return input value.
	return value.(bool), nil
}
