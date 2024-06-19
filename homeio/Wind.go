package homeio

import "fmt"

func (h *home) GetWindSpeed() (float64, error) {
	value, err := h.values["wdsp"]
	if !err {
		return 0, fmt.Errorf("wind speed sensor not found")
	}
	return value.(float64), nil
}
