package nrgsim

// Storable is an interface that defines the methods that need to be implemented by objects that can store energy. The GetCapacity method should return the total capacity of the
// storage in Joules, the GetStoredEnergy method should return the total stored energy in Joules, and the GetStateOfCharge method should return the state of charge of the storage
// as a percentage.
type Storable interface {
	Measurable

	// GetCapacity returns the total capacity of the storage in Joules.
	GetCapacity() float64

	// GetStoredEnergy returns the total stored energy in Joules.
	GetStoredEnergy() float64

	// GetStateOfCharge returns the state of charge of the storage as a percentage.
	GetStateOfCharge() float64
}
