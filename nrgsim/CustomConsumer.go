package nrgsim

import "time"

type (
	// PowerReadFunction is a callback function that returns the current power consumption of a customConsumer. It will be called by the customConsumer at every simulation step, so it should be
	// implemented efficiently.
	PowerReadFunction func() float64

	customConsumer struct {
		baseMeasurable
		powerReadFunction PowerReadFunction
	}

	// CustomConsumer is the interface that can be used to control a customConsumer in the simulation.
	CustomConsumer interface {
		Measurable

		// GetNominalPower returns the nominal power of the customConsumer.
		GetNominalPower() float64

		// SetNominalPower allows to change the nominal power of the customConsumer at any time.
		SetNominalPower(nominalPower float64)
	}

	customConsumerObject struct {
		customConsumer
		nominalPower float64
	}
)

/* simulable interface implementation for customConsumer */

func (c *customConsumer) calculateStep(deltaTime time.Duration) {
	// Read the current power consumption from the power read function.
	c.power = c.powerReadFunction()

	// Update the energy consumption of the customConsumer.
	c.energy += c.power * deltaTime.Seconds()
}

/* CustomConsumer interface implementation for customConsumerObject */

func (c *customConsumerObject) GetNominalPower() float64 {
	return c.nominalPower
}

func (c *customConsumerObject) SetNominalPower(nominalPower float64) {
	c.nominalPower = nominalPower
}

func (c *customConsumerObject) getPower() float64 {
	return c.nominalPower
}
