package homeio

import "fmt"

func (h *home) IsSmokeDetected(room Room) (bool, error) {
	value, ok := h.values[fmt.Sprintf("sdtc/%s", room)]
	if !ok {
		return false, fmt.Errorf("smoke detector not found")
	}
	return value.(bool), nil
}
