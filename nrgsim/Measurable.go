package nrgsim

// Measurable is an interface that defines the methods that need to be implemented by objects that need to be measured. The GetPower method should return the current power consumption
// of the object in Watts, the GetEnergy method should return the total energy consumption of the object in Joules, and the Reset method should reset the energy consumption
// of the object to zero.
type Measurable interface {
	// GetPower returns the current power consumption of the object in Watts.
	GetPower() float64

	// GetEnergy returns the total energy consumption of the object in Joules.
	GetEnergy() float64

	// Reset resets the energy consumption of the object to zero.
	Reset()
}
