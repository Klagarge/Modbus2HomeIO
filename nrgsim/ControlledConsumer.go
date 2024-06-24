package nrgsim

import (
	"time"
)

type (
	// ControlFunction is a function that returns a boolean value to indicate whether the consumer is on or off. The consumer is on if the function returns true. The function will
	// be called by the	consumer at every simulation step, so it should be implemented efficiently.
	ControlFunction func() bool

	controlledConsumer struct {
		baseMeasurable

		controlFunction ControlFunction
		nominalPower    float64
	}

	// ControlledConsumer is the interface that can be used to control a controller consumer object in the simulation.
	ControlledConsumer interface {
		Measurable

		// GetNominalPower returns the nominal power of the controlled consumer.
		GetNominalPower() float64

		// SetNominalPower allows to change the nominal power of the controlled consumer at any time.
		SetNominalPower(nominalPower float64)

		// IsOn returns true if the controlled consumer is turned on, false otherwise.
		IsOn() bool

		// SetOn allows to turn the controlled consumer on or off at any time.
		SetOn(on bool)
	}

	controlledConsumerObject struct {
		controlledConsumer
		nominalPower float64
		isOn         bool
	}
)

/* simulable interface implementation for controlledConsumer */

func (c *controlledConsumer) calculateStep(duration time.Duration) {
	// Read the current state from the control function.
	on := c.controlFunction()

	// Update the power consumption of the controlled consumer.
	c.power = map[bool]float64{true: c.nominalPower, false: 0}[on]

	// Update the energy consumption of the controlled consumer.
	c.energy += c.power * duration.Seconds()
}

/* ControlledConsumer interface implementation for controlledConsumerObject */

func (c *controlledConsumerObject) GetNominalPower() float64 {
	return c.nominalPower
}

func (c *controlledConsumerObject) SetNominalPower(nominalPower float64) {
	c.nominalPower = nominalPower
}

func (c *controlledConsumerObject) IsOn() bool {
	return c.isOn
}

func (c *controlledConsumerObject) SetOn(on bool) {
	c.isOn = on
}
