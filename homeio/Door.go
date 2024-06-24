package homeio

import "fmt"

// Door identifies a door in a room. Up to 2 doors are supported in a room. A door can be either a normal door or a window that can be opened.
type Door int

const (
	// Door1 is the first door in a room.
	Door1 Door = 1

	// Door2 is the second door in a room.
	Door2 Door = 2
)

func (h *home) IsDoorClosed(room Room, door Door) (bool, error) {
	// Read input value from inputs map and check if it exists.
	value, ok := h.inputs[fmt.Sprintf("ddtc/%d/%s", door, room)]
	if !ok {
		return false, fmt.Errorf("door sensor not found")
	}

	// Return input value.
	return value.(bool), nil
}
