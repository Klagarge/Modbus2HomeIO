package homeio

import "fmt"

func (h *home) IsMotionDetected(room Room) (bool, error) {
	value, ok := h.values[fmt.Sprintf("mdtc/%s", room)]
	if !ok {
		return false, fmt.Errorf("motion sensor not found")
	}
	return value.(bool), nil
}
