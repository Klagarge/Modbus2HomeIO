package homeio

import "fmt"

// Door identifies a door in the house.
type Door int

const (
	Door1 Door = 1
	Door2 Door = 2
)

func (h *home) IsDoorOpen(room Room, door Door) (bool, error) {
	value, ok := h.values[fmt.Sprintf("ddtc/%d/%s", door, room)]
	if !ok {
		return false, fmt.Errorf("door sensor not found")
	}
	return value.(bool), nil
}
