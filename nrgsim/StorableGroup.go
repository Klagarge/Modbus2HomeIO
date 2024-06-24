package nrgsim

// StorableGroup is a group of Storable objects that can be treated as a single Storable object. The GetPower method returns the total power consumption of the group in Watts,
// the GetEnergy method returns the total energy consumption of the group in Joules, the Reset method resets the energy consumption of all the objects in the group to zero,
// the GetStateOfCharge method returns the state of charge of the group as a percentage, and the GetCapacity method returns the total capacity of the group in Joules.
type StorableGroup map[string]Storable

/* Storable interface implementation */

func (s *StorableGroup) GetPower() float64 {
	// Power is just the sum of the power of all the storables in the group.
	power := 0.0
	for _, consumer := range *s {
		power += consumer.GetPower()
	}
	return power
}

func (s *StorableGroup) GetEnergy() float64 {
	// Consumed/Produced Energy is just the sum of the energy of all the storables in the group.
	energy := 0.0
	for _, consumer := range *s {
		energy += consumer.GetEnergy()
	}
	return energy
}

func (s *StorableGroup) Reset() {
	// Reset the energy consumption of all the storables in the group.
	for _, consumer := range *s {
		consumer.Reset()
	}
}

func (s *StorableGroup) GetStateOfCharge() float64 {
	// State of charge is the average state of charge of all the storables in the group weighted by the storage's capacity.
	stateOfCharge := 0.0
	totalCapacity := 0.0
	for _, consumer := range *s {
		stateOfCharge += consumer.GetStateOfCharge() * consumer.GetCapacity()
		totalCapacity += consumer.GetCapacity()
	}
	return stateOfCharge / totalCapacity
}

func (s *StorableGroup) GetCapacity() float64 {
	// Capacity is just the sum of the capacity of all the storables in the group.
	capacity := 0.0
	for _, consumer := range *s {
		capacity += consumer.GetCapacity()
	}
	return capacity
}
