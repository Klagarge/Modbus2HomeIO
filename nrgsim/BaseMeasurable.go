package nrgsim

// baseMeasurable is a struct that implements the Measurable interface. It is used as a base for other structs that need to implement the Measurable interface.
type baseMeasurable struct {
	// power is the current power consumption of the consumer in Watts.
	power float64

	// energy is the total energy consumption of the consumer in Joules.
	energy float64
}

/* Measurable interface implementation for baseMeasurable */

func (b *baseMeasurable) GetPower() float64 {
	return b.power
}

func (b *baseMeasurable) GetEnergy() float64 {
	return b.energy
}

func (b *baseMeasurable) Reset() {
	b.power = 0.0
	b.energy = 0.0
}
