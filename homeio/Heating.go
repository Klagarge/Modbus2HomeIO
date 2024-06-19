package homeio

import (
	"fmt"
	"net/http"
)

func (h *home) SetHeatingOn(room Room, on bool) error {
	response, err := h.client.Get(fmt.Sprintf("%s/swh/%s/%s", h.baseURL, map[bool]string{true: "turn_on", false: "turn_off"}[on], room))
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", response.Status)
	}
	return nil
}

func (h *home) SetHeatingLevel(room Room, level float64) error {
	response, err := h.client.Get(fmt.Sprintf("%s/sth/%f/%s", h.baseURL, level*10.0, room))
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", response.Status)
	}
	return nil
}
