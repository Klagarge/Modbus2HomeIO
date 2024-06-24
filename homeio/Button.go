package homeio

import "fmt"

// Button identifies a button in a simulated room. Up to 5 buttons are supported in a room.
type Button int

const (
	// Button1 is the first button in a room.
	Button1 Button = 1

	// Button2 is the second button in a room.
	Button2 Button = 2

	// Button3 is the third button in a room.
	Button3 Button = 3

	// Button4 is the fourth button in a room.
	Button4 Button = 4

	// Button5 is the fifth button in a room.
	Button5 Button = 5
)

/* Home interface implementation */

func (h *home) IsButtonPressed(room Room, button Button) (bool, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("lsw/%d/%s", button, room)]
	if !ok {
		return false, fmt.Errorf("button not found")
	}

	// Return input value.
	return value.(bool), nil
}
