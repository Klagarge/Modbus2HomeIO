package homeio

import "fmt"

func (h *home) GetRelativeHumidity() (float64, error) {
	value, err := h.values["rhm"]
	if !err {
		return 0, fmt.Errorf("relative humidity sensor not found")
	}
	return value.(float64), nil
}
