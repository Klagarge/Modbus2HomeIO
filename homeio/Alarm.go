package homeio

import (
	"fmt"
	"net/http"
)

func (h *home) IsAlarmActive() (bool, error) {
	value, ok := h.values["aa"]
	if !ok {
		return false, fmt.Errorf("alarm not found")
	}
	return value.(bool), nil
}

func (h *home) SetAlarmOn(on bool) error {
	response, err := h.client.Get(fmt.Sprintf("%s/sal/%s", h.baseURL, map[bool]string{true: "turn_on", false: "turn_off"}[on]))
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", response.Status)
	}
	return nil
}
