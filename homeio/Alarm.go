package homeio

import (
	"fmt"
	"net/http"
)

/* Home interface implementation */

func (h *home) IsAlarmArmed() (bool, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs["aa"]
	if !ok {
		return false, fmt.Errorf("alarm not found")
	}

	// Return input value.
	return value.(bool), nil
}

func (h *home) SetAlarmArmed(armed bool) error {
	// Set local state in outputs map.
	h.outputs["sal"] = armed

	// Send request and check response status.
	response, err := h.client.Get(fmt.Sprintf("%s/sal/%s", h.baseURL, map[bool]string{true: "turn_on", false: "turn_off"}[armed]))
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", response.Status)
	}

	// Return nil if everything went well.
	return nil
}

func (h *home) GetAlarmArmedOutput() (bool, error) {
	// Read output value from outputs map and check if it exists.
	value, ok := h.outputs["sal"]
	if !ok {
		return false, fmt.Errorf("alarm not found")
	}

	// Return output value.
	return value.(bool), nil
}
