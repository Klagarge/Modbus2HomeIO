package homeio

import (
	"time"
)

type (
	Home interface {
		Poll() error

		IsButtonPressed(room Room, button Button) (bool, error)
		IsDimmerControlUpPressed(room Room, button Dimmer) (bool, error)
		IsDimmerControlDownPressed(room Room, button Dimmer) (bool, error)
		IsShutterControlUpPressed(room Room, shutter Shutters) (bool, error)
		IsShutterControlDownPressed(room Room, shutter Shutters) (bool, error)
		IsDoorOpen(room Room, door Door) (bool, error)
		IsSmokeDetected(room Room) (bool, error)
		IsMotionDetected(room Room) (bool, error)
		IsLightDetected(room Room) (bool, error)
		IsGateOpen(room Room) (bool, error)
		IsGateClosed(room Room) (bool, error)
		IsIRDetected(room Room, sensor IRSensor) (bool, error)
		IsRemoteControlButtonPressed(button RemoteControlButton) (bool, error)
		IsAlarmActive() (bool, error)

		GetShutterPosition(room Room, shutter Shutters) (uint16, error)
		GetBrightness(room Room) (uint16, error)
		GetTemperature(room Room) (uint16, error)
		GetTemperateSetPoint(room Room) (float64, error)
		GetOutsideTemperature() (float64, error)
		GetRelativeHumidity() (float64, error)
		GetWindSpeed() (float64, error)
		GetPosition() (float64, float64, error)
		GetTime() (time.Time, error)

		SetLightOn(room Room, light Light, on bool) error
		SetLightDim(room Room, light Light, dim uint16) error
		SetShuttersDirection(room Room, shutter Shutters, direction ShuttersDirection) error
		SetGateDirection(room Room, direction GateDirection) error
		SetHeatingOn(room Room, on bool) error
		SetHeatingLevel(room Room, level float64) error
		SetAlarmOn(on bool) error
	}
)
