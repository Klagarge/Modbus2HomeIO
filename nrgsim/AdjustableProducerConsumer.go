package nrgsim

import (
	"time"
)

type (
	// AdjustFunction is a function that returns a float64 value to adjust the power of an adjustableProducerConsumer. It han have a value between 0.0 and 1.0 to scale set the
	// power to a fraction of the nominal power, or a value greater than 1 is supported to simulate power on peak power for consumers. Even a negative factor is supported allowing
	// to simulate power generation of systems with great inertia powered off. It will be called by the adjustableProducerConsumer at every simulation step, so it should be
	// implemented efficiently.
	AdjustFunction func() float64

	adjustableProducerConsumer struct {
		baseMeasurable

		isProducer     bool
		adjustFunction AdjustFunction
		nominalPower   float64
	}

	AdjustableProducer interface {
		Measurable

		// GetNominalPower returns the nominal power of the adjustable producer or consumer.
		GetNominalPower() float64

		// SetNominalPower allows to change the nominal power of the adjustable producer or consumer at any time.
		SetNominalPower(nominalPower float64)

		// GetFactor returns the level of the adjustable producer or consumer.
		GetFactor() float64

		// SetFactor allows to change the level of the adjustable producer or consumer at any time.
		SetFactor(level float64)
	}

	AdjustableConsumer interface {
		AdjustableProducer
	}

	adjustableProducerConsumerObject struct {
		adjustableProducerConsumer
		factor float64
	}
)

/* simulable interface implementation for adjustableProducerConsumer */

func (a *adjustableProducerConsumer) calculateStep(duration time.Duration) {
	// Calculate the power factor.
	factor := a.adjustFunction()

	// Update the power and energy consumption of the adjustableProducerConsumer.
	a.power = a.nominalPower * factor * map[bool]float64{true: -1, false: 1}[a.isProducer]

	// Update the energy consumption of the adjustableProducerConsumer.
	a.energy += a.power * duration.Seconds()
}

/* AdjustableProducer/AdjustableConsumer interface implementation for adjustableProducerConsumerObject */

func (a *adjustableProducerConsumerObject) GetNominalPower() float64 {
	return a.nominalPower
}

func (a *adjustableProducerConsumerObject) SetNominalPower(nominalPower float64) {
	a.nominalPower = nominalPower
}

func (a *adjustableProducerConsumerObject) GetFactor() float64 {
	return a.factor
}

func (a *adjustableProducerConsumerObject) SetFactor(factor float64) {
	a.factor = factor
}
