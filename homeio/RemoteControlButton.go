package homeio

import "fmt"

// RemoteControlButton identifies a button on the virtual remote control.
type RemoteControlButton int

const (
	// RemoteControlButton1 is the first button on the remote control.
	RemoteControlButton1 RemoteControlButton = 1

	// RemoteControlButton2 is the second button on the remote control.
	RemoteControlButton2 RemoteControlButton = 2

	// RemoteControlButton3 is the third button on the remote control.
	RemoteControlButton3 RemoteControlButton = 3

	// RemoteControlButton4 is the fourth button on the remote control.
	RemoteControlButton4 RemoteControlButton = 4

	// RemoteControlButton5 is the fifth button on the remote control.
	RemoteControlButton5 RemoteControlButton = 5

	// RemoteControlButton6 is the sixth button on the remote control.
	RemoteControlButton6 RemoteControlButton = 6

	// RemoteControlButton7 is the seventh button on the remote control.
	RemoteControlButton7 RemoteControlButton = 7

	// RemoteControlButton8 is the eighth button on the remote control.
	RemoteControlButton8 RemoteControlButton = 8
)

/* Home interface implementation */

func (h *home) IsRemoteControlButtonPressed(button RemoteControlButton) (bool, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("rmt/%d", button)]
	if !ok {
		return false, fmt.Errorf("remote button not found")
	}

	// Return input value.
	return value.(bool), nil
}
