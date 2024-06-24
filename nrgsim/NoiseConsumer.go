package nrgsim

import (
	"math/rand"
	"time"
)

type (
	noiseConsumer struct {
		baseMeasurable

		minimumPower float64
		maximumPower float64
		interval     time.Duration
		remaining    time.Duration
	}

	// NoiseConsumer is the interface that can be used to control a noise consumer in the simulation.
	NoiseConsumer interface {
		Measurable

		// GetPowerLevels returns the minimum and maximum power levels of the noise consumer.
		GetPowerLevels() (float64, float64)

		// SetPowerLevels allows to change the minimum and maximum power levels of the noise consumer at any time.
		SetPowerLevels(minimumPower float64, maximumPower float64)

		// GetInterval returns the interval at which the power level of the noise consumer changes.
		GetInterval() time.Duration

		// SetInterval allows to change the interval at which the power level of the noise consumer changes at any time.
		SetInterval(interval time.Duration)
	}

	noiseConsumerObject struct {
		noiseConsumer
	}
)

/* simulation interface implementation for noiseConsumer */

func (n *noiseConsumer) calculateStep(duration time.Duration) {
	// Subtract the duration from the remaining time.
	n.remaining -= duration

	// If the remaining time is less than or equal to zero, generate a new power level and reset the remaining time.
	if n.remaining.Seconds() <= 0 {
		// Reset the remaining time.
		n.remaining = n.interval

		// Generate a new random power level.
		n.power = n.minimumPower + (n.maximumPower-n.minimumPower)*rand.Float64()

		// Calculate the energy.
		n.energy += n.power * duration.Seconds()
	}
}

/* NoiseConsumer interface implementation for noiseConsumerObject */

func (n *noiseConsumerObject) GetPowerLevels() (float64, float64) {
	return n.minimumPower, n.maximumPower
}

func (n *noiseConsumerObject) SetPowerLevels(minimumPower float64, maximumPower float64) {
	n.minimumPower = minimumPower
	n.maximumPower = maximumPower
}

func (n *noiseConsumerObject) GetInterval() time.Duration {
	return n.interval
}

func (n *noiseConsumerObject) SetInterval(interval time.Duration) {
	n.interval = interval
	n.remaining = interval
}
