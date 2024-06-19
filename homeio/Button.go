package homeio

import "fmt"

// Button identifies a switch in the house.
type Button int

const (
	Button1 Button = 1
	Button2 Button = 2
	Button3 Button = 3
	Button4 Button = 4
	Button5 Button = 5
)

func (h *home) IsButtonPressed(room Room, button Button) (bool, error) {
	value, ok := h.values[fmt.Sprintf("lsw/%d/%s", button, room)]
	if !ok {
		return false, fmt.Errorf("switch not found")
	}
	return value.(bool), nil
}
