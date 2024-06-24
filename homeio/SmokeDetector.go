package homeio

import "fmt"

/* Home interface implementation */

func (h *home) IsSmokeDetected(room Room) (bool, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("sdtc/%s", room)]
	if !ok {
		return false, fmt.Errorf("smoke detector not found")
	}

	// Return input value.
	return value.(bool), nil
}
