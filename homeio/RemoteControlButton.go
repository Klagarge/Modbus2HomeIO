package homeio

import "fmt"

// RemoteControlButton identifies a button on the remote control.
type RemoteControlButton int

const (
	RemoteControlButton1 RemoteControlButton = 1
	RemoteControlButton2 RemoteControlButton = 2
	RemoteControlButton3 RemoteControlButton = 3
	RemoteControlButton4 RemoteControlButton = 4
	RemoteControlButton5 RemoteControlButton = 5
	RemoteControlButton6 RemoteControlButton = 6
	RemoteControlButton7 RemoteControlButton = 7
	RemoteControlButton8 RemoteControlButton = 8
)

func (h *home) IsRemoteControlButtonPressed(button RemoteControlButton) (bool, error) {
	value, ok := h.values[fmt.Sprintf("rmt/%d", button)]
	if !ok {
		return false, fmt.Errorf("remote button not found")
	}
	return value.(bool), nil
}
