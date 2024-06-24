package homeio

import (
	"time"
)

// Home defines the interface for the Home I/O REST API to get sensor data and control actuators in the game.
type Home interface {

	// Poll Gets the latest sensor data from the Home I/O REST API and caches it to be read by the getter methods.
	Poll() error

	/* Digital Inputs */

	// IsButtonPressed Gets the state of a button. Returns true if the button is pressed, false otherwise. Returns an error if the button state cannot be read. The room parameter
	// specifies the room in which the button is located. The button parameter specifies the button number to read.
	IsButtonPressed(room Room, button Button) (bool, error)

	// IsDimmerControlUpPressed Gets the state of a dimmer control up button. Returns true if the button is pressed, false otherwise. Returns an error if the button state cannot
	// be read. The room parameter specifies the room in which the dimmer control is located. The button parameter specifies the dimmer control button number to read.
	IsDimmerControlUpPressed(room Room, button Dimmer) (bool, error)

	// IsDimmerControlDownPressed Gets the state of a dimmer control down button. Returns true if the button is pressed, false otherwise. Returns an error if the button state
	// cannot be read. The room parameter specifies the room in which the dimmer control is located. The button parameter specifies the dimmer control button number to read.
	IsDimmerControlDownPressed(room Room, button Dimmer) (bool, error)

	// IsShutterControlUpPressed Gets the state of a shutter control up button. Returns true if the button is pressed, false otherwise. Returns an error if the button state cannot
	// be read. The room parameter specifies the room in which the shutter control is located. The shutter parameter specifies the shutter control button number to read.
	IsShutterControlUpPressed(room Room, shutter Shutters) (bool, error)

	// IsShutterControlDownPressed Gets the state of a shutter control down button. Returns true if the button is pressed, false otherwise. Returns an error if the button state
	// cannot be read. The room parameter specifies the room in which the shutter control is located. The shutter parameter specifies the shutter control button number to read.
	IsShutterControlDownPressed(room Room, shutter Shutters) (bool, error)

	// AreShuttersOnTop Gets the state of a shutter. Returns true if the shutter is on top, false otherwise. Returns an error if the shutter state cannot be read. The room
	// parameter specifies the room in which the shutter is located. The shutter parameter specifies the shutter number to read.
	AreShuttersOnTop(room Room, shutter Shutters) (bool, error)

	// AreShuttersOnBottom Gets the state of a shutter. Returns true if the shutter is on bottom, false otherwise. Returns an error if the shutter state cannot be read. The room
	// parameter specifies the room in which the shutter is located. The shutter parameter specifies the shutter number to read.
	AreShuttersOnBottom(room Room, shutter Shutters) (bool, error)

	// IsDoorClosed Gets the state of a door. Returns true if the door is closed, false otherwise. Returns an error if the door state cannot be read. The room parameter specifies
	// the room in which the door is located. The door parameter specifies the door number to read.
	IsDoorClosed(room Room, door Door) (bool, error)

	// IsSmokeDetected Gets the state of a smoke detector. Returns true if smoke is detected, false otherwise. Returns an error if the smoke detector state cannot be read. The room
	// parameter specifies the room in which the smoke detector is located.
	IsSmokeDetected(room Room) (bool, error)

	// IsMotionDetected Gets the state of a motion detector. Returns true if motion is detected, false otherwise. Returns an error if the motion detector state cannot be read. The
	// room parameter specifies the room in which the motion detector is located.
	IsMotionDetected(room Room) (bool, error)

	// IsLightDetected Gets the state of a light detector. Returns true if light is detected, false otherwise. Returns an error if the light detector state cannot be read. The room
	// parameter specifies the room in which the light detector is located. DOES NOT WORK CURRENTLY.
	IsLightDetected(room Room) (bool, error)

	// IsGateOpen Gets the state of a gate. Returns true if the gate is open, false otherwise. Returns an error if the gate state cannot be read. The room parameter specifies the
	// room in which the gate is located. If the room is Exterior, it represents the main gate. If the room is Garage, the garage door is represented.
	IsGateOpen(room Room) (bool, error)

	// IsGateClosed Gets the state of a gate. Returns true if the gate is closed, false otherwise. Returns an error if the gate state cannot be read. The room parameter specifies the
	// room in which the gate is located. If the room is Exterior, the gate is the main gate. If the room is Garage, the gate is the garage door.
	IsGateClosed(room Room) (bool, error)

	// IsIRObscured Gets the state of an infrared sensor. Returns true if the sensor is obscured, false otherwise. Returns an error if the sensor state cannot be read. The room
	// parameter specifies the room in which the sensor is located. The sensor parameter specifies the infrared sensor number to read. Only the Exterior room (3) and the Garage
	// room (1) have infrared sensors.
	IsIRObscured(room Room, sensor IRSensor) (bool, error)

	// IsRemoteControlButtonPressed Gets the state of a remote control button. Returns true if the button is pressed, false otherwise. Returns an error if the button state cannot be
	// read. The button parameter specifies the remote control button number to read. The game disposes 8 remote control buttons.
	IsRemoteControlButtonPressed(button RemoteControlButton) (bool, error)

	// IsAlarmArmed Gets the state of the alarm system. Returns true if the alarm is armed, false otherwise. Returns an error if the alarm state cannot be read.
	IsAlarmArmed() (bool, error)

	/* Analog Inputs */

	// GetShutterPosition gets the position of a shutter. Returns the shutter position in per thousand. Returns an error if the shutter position cannot be read. The room parameter
	// specifies the room in which the shutter is located. The shutter parameter specifies the shutter number to read.
	GetShutterPosition(room Room, shutter Shutters) (uint16, error)

	// GetBrightness gets the brightness level in percent to a maximum. Returns an error if the brightness level cannot be read. The room parameter specifies the room in which the
	// brightness sensor is located.
	GetBrightness(room Room) (uint16, error)

	// GetTemperature gets the temperature in degrees Celsius. Returns an error if the temperature cannot be read. The room parameter specifies the room in which the temperature
	// sensor is located.
	GetTemperature(room Room) (uint16, error)

	// GetTemperateSetPoint gets a thermostat's temperature set point in degrees Celsius. Returns an error if the temperature set point cannot be read. The room parameter specifies
	// the room in which the thermostat is located.
	GetTemperateSetPoint(room Room) (uint16, error)

	// GetOutsideTemperature gets the outside temperature in Kelvin. Returns an error if the outside temperature cannot be read.
	GetOutsideTemperature() (uint16, error)

	// GetRelativeHumidity gets the relative humidity in percent. Returns an error if the relative humidity cannot be read.
	GetRelativeHumidity() (uint16, error)

	// GetWindSpeed gets the wind speed in meters per second. Returns an error if the wind speed cannot be read.
	GetWindSpeed() (uint16, error)

	// GetPosition gets the position of the building in latitude and longitude. Returns the latitude and longitude in degrees. Returns an error if the position cannot be read.
	GetPosition() (float64, float64, error)

	// GetTime gets the current time in the simulation. Returns the current time. Returns an error if the time cannot be read.
	GetTime() (time.Time, error)

	// GetYear gets the current year in the simulation. Returns the current year. Returns an error if the year cannot be read.
	GetYear() (uint16, error)

	// GetMonth gets the current month in the simulation. Returns the current month. Returns an error if the month cannot be read.
	GetMonth() (uint16, error)

	// GetDay gets the current day in the simulation. Returns the current day. Returns an error if the day cannot be read.
	GetDay() (uint16, error)

	// GetHour gets the current hour in the simulation. Returns the current hour. Returns an error if the hour cannot be read.
	GetHour() (uint16, error)

	// GetMinute gets the current minute in the simulation. Returns the current minute. Returns an error if the minute cannot be read.
	GetMinute() (uint16, error)

	// GetSecond gets the current second in the simulation. Returns the current second. Returns an error if the second cannot be read.
	GetSecond() (uint16, error)

	/* Outputs */

	// SetLightOn sets the state of a light. The on parameter specifies whether the light should be on or off. Returns an error if the light state cannot be set. The room parameter
	// specifies the room in which the light is located. The light parameter specifies the light number to control.
	SetLightOn(room Room, light Light, on bool) error

	// SetDimmerPercentage sets the dimmer percentage of a light. The dim parameter specifies the dimmer percentage in percent. Returns an error if the dimmer percentage cannot
	// be set. The room parameter specifies the room in which the light is located. The light parameter specifies the light number to control.
	SetDimmerPercentage(room Room, light Light, dim uint16) error

	// GetDimmerPercentageOutput gets the dimmer percentage of a light. Returns the dimmer percentage in percent. Returns an error if the dimmer percentage cannot be read. The room
	// parameter specifies the room in which the light is located. The light parameter specifies the light number to read.
	GetDimmerPercentageOutput(room Room, light Light) (uint16, error)

	// SetShuttersDirection sets the direction of a shutter. The direction parameter specifies the shutter direction. Returns an error if the shutter direction cannot be set. The
	// room parameter specifies the room in which the shutter is located. The shutter parameter specifies the shutter number to control.
	SetShuttersDirection(room Room, shutter Shutters, direction ShuttersDirection) error

	// GetShuttersDirectionOutput gets the direction of a shutter. Returns the shutter direction. Returns an error if the shutter direction cannot be read. The room parameter
	// specifies the room in which the shutter is located. The shutter parameter specifies the shutter number to read.
	GetShuttersDirectionOutput(room Room, shutter Shutters) (ShuttersDirection, error)

	// SetGateDirection sets the direction of a gate. The direction parameter specifies the gate direction. Returns an error if the gate direction cannot be set. The room parameter
	// specifies the room in which the gate is located. If the room is Exterior, the gate is the main gate. If the room is Garage, the gate is the garage door.
	SetGateDirection(room Room, direction GateDirection) error

	// GetGateDirectionOutput gets the direction of a gate. Returns the gate direction. Returns an error if the gate direction cannot be read. The room parameter specifies the room
	// in which the gate is located. If the room is Exterior, the gate is the main gate. If the room is Garage, the gate is the garage door.
	GetGateDirectionOutput(room Room) (GateDirection, error)

	// SetHeatingOn sets the state of a heating system. The on parameter specifies whether the heating system should be on or off. Returns an error if the heating system state
	// cannot be set. The room parameter specifies the room in which the heating system is located.
	SetHeatingOn(room Room, on bool) error

	// SetHeatingPower sets the heating power of a heating system. The power parameter specifies the heating power in watts. Returns an error if the heating power cannot be set.
	// The room parameter specifies the room in which the heating system is located.
	SetHeatingPower(room Room, power uint16) error

	// GetHeatingPowerOutput gets the heating power of a heating system. Returns the heating power in watts. Returns an error if the heating power cannot be read. The room
	// parameter specifies the room in which the heating system is located.
	GetHeatingPowerOutput(room Room) (uint16, error)

	// SetSirenOn sets the state of a siren. The on parameter specifies whether the siren should be on or off. Returns an error if the siren state cannot be set. The room parameter
	// specifies the room in which the siren is located. Sirens are installed in the Exterior room and the EntryHall room.
	SetSirenOn(room Room, on bool) error

	// GetSirenOnOutput gets the state of a siren. Returns true if the siren is on, false otherwise. Returns an error if the siren state cannot be read. The room parameter specifies
	// the room in which the siren is located. Sirens are installed in the Exterior room and the EntryHall room.
	GetSirenOnOutput(room Room) (bool, error)

	// SetAlarmArmed sets the state of the alarm system. The on parameter specifies whether the alarm system should be armed or disarmed. Returns an error if the alarm state cannot
	// be set.
	SetAlarmArmed(on bool) error

	// GetAlarmArmedOutput gets the state of the alarm system. Returns true if the alarm is set to be armed, false otherwise. Returns an error if the alarm state cannot be read.
	GetAlarmArmedOutput() (bool, error)
}
