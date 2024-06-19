package homeio

import "fmt"

func (h *home) GetPosition() (float64, float64, error) {
	lat, err := h.values["lat"]
	if !err {
		return 0, 0, fmt.Errorf("latitude not found")
	}
	long, err := h.values["long"]
	if !err {
		return 0, 0, fmt.Errorf("longitude not found")
	}
	return lat.(float64), long.(float64), nil
}
