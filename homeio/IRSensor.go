package homeio

import "fmt"

// IRSensor identifies an infrared sensor in a room. Up to 3 infrared sensors are supported in the main gate and 1 in the garage.
type IRSensor int

const (
	// EntranceIRSensor1 is the first infrared sensor of the main gate.
	EntranceIRSensor1 IRSensor = 1

	// EntranceIRSensor2 is the second infrared sensor of the main gate.
	EntranceIRSensor2 IRSensor = 2

	// EntranceIRSensor3 is the third infrared sensor of the main gate.
	EntranceIRSensor3 IRSensor = 3

	// GarageIRSensor is the infrared sensor of the garage.
	GarageIRSensor IRSensor = 1
)

/* Home interface implementation */

func (h *home) IsIRObscured(room Room, sensor IRSensor) (bool, error) {
	switch room {
	case Exterior:
		// Check if the sensor number is valid.
		if sensor < EntranceIRSensor1 || sensor > EntranceIRSensor3 {
			return false, fmt.Errorf("invalid sensor number")
		}

		// Read input value from inputs map and check if it exists.
		value, ok := h.inputs[fmt.Sprintf("gtde/infrared_%d", sensor)]
		if !ok {
			return false, fmt.Errorf("infrared sensor not found")
		}

		// Return input value.
		return value.(bool), nil

	case Garage:
		// Check if the sensor number is valid.
		if sensor != GarageIRSensor {
			return false, fmt.Errorf("invalid sensor number")
		}

		// Read input value from inputs map and check if it exists.
		value, ok := h.inputs[fmt.Sprintf("gtdg")]
		if !ok {
			return false, fmt.Errorf("infrared sensor not found")
		}

		// Return input value.
		return value.(bool), nil

	default:
		// Return error if the room is invalid.
		return false, fmt.Errorf("invalid room")
	}
}
