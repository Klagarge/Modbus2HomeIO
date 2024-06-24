package homeio

import (
	"fmt"
	"math"
	"net/http"
)

// heatingMaxPowers defines the maximum power of the electrical heater for each room that disposes of such.
var heatingMaxPowers = map[Room]float64{
	LivingRoom:      2000,
	Kitchen:         1500,
	EntranceHall:    1750,
	BedroomCorridor: 1500,
	ChildrenRoom:    1000,
	Bathroom:        750,
	SingleBedroom:   750,
	PrivateBathroom: 1000,
	CoupleBedroom:   750,
	LaundryRoom:     500,
	HomeOffice:      1500,
}

/* Home interface implementation */

func (h *home) SetHeatingOn(room Room, on bool) error {
	// Get the maximum power for the room and check if it exists.
	maxPower, ok := heatingMaxPowers[room]
	if !ok {
		return fmt.Errorf("invalid room")
	}

	// Set local state in outputs map.
	if on {
		h.outputs[fmt.Sprintf("sth/%s", room)] = maxPower
	} else {
		h.outputs[fmt.Sprintf("sth/%s", room)] = 0
	}

	// Send request and check response status.
	response, err := h.client.Get(fmt.Sprintf("%s/swh/%s/%s", h.baseURL, map[bool]string{true: "turn_on", false: "turn_off"}[on], room))
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", response.Status)
	}

	// Return nil if everything went well.
	return nil
}

func (h *home) SetHeatingPower(room Room, power uint16) error {
	// Get the maximum power for the room and check if it exists.
	maxPower, ok := heatingMaxPowers[room]
	if !ok {
		return fmt.Errorf("invalid room")
	}

	// Limit the power to the maximum power of the room.
	effectivePower := math.Min(maxPower, float64(power))

	// Set local state in outputs map.
	h.outputs[fmt.Sprintf("sth/%s", room)] = effectivePower
	response, err := h.client.Get(fmt.Sprintf("%s/sth/%f/%s", h.baseURL, float64(power)/maxPower*10.0, room))
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response status: %s", response.Status)
	}

	// Return nil if everything went well.
	return nil
}

func (h *home) GetHeatingPowerOutput(room Room) (uint16, error) {
	// Get output value from outputs map and check if it exists.
	value, ok := h.outputs[fmt.Sprintf("sth/%s", room)]
	if !ok {
		return 0, fmt.Errorf("heating not found")
	}

	// Return output value.
	return uint16(value.(float64)), nil
}
