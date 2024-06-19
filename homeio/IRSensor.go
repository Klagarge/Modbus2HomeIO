package homeio

import "fmt"

// IRSensor identifies an infrared sensor in the house.
type IRSensor int

const (
	EntranceIRSensor1 IRSensor = 1
	EntranceIRSensor2 IRSensor = 2
	EntranceIRSensor3 IRSensor = 3
	GarageIRSensor    IRSensor = 1
)

func (h *home) IsIRDetected(room Room, sensor IRSensor) (bool, error) {
	switch room {
	case Exterior:
		value, ok := h.values[fmt.Sprintf("gtde/infrared_%d", sensor)]
		if !ok {
			return false, fmt.Errorf("infrared sensor not found")
		}
		return value.(bool), nil

	case Garage:
		if sensor != GarageIRSensor {
			return false, fmt.Errorf("invalid sensor number")
		}
		value, ok := h.values[fmt.Sprintf("gtdg")]
		if !ok {
			return false, fmt.Errorf("infrared sensor not found")
		}
		return value.(bool), nil

	default:
		return false, fmt.Errorf("invalid room")
	}
}
