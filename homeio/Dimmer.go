package homeio

import "fmt"

// Dimmer identifies a dimmer in the house.
type Dimmer int

const (
	Dimmer1 Dimmer = 1
	Dimmer2 Dimmer = 2
	Dimmer3 Dimmer = 3
)

func (h *home) IsDimmerControlUpPressed(room Room, dimmer Dimmer) (bool, error) {
	value, ok := h.values[fmt.Sprintf("lsd/%d/%s/up", dimmer, room)]
	if !ok {
		return false, fmt.Errorf("dimmer not found")
	}
	return value.(bool), nil
}

func (h *home) IsDimmerControlDownPressed(room Room, dimmer Dimmer) (bool, error) {
	value, ok := h.values[fmt.Sprintf("lsd/%d/%s/down", dimmer, room)]
	if !ok {
		return false, fmt.Errorf("dimmer not found")
	}
	return value.(bool), nil
}
