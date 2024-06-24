package nrgsim

// MeasurableGroup is a named collection of objects implementing the Measurable interfaces. It is used to group multiple Measurables together and provide a simple way to
// calculate the total power and energy consumption of all the objects in the group by implementing the Measurable interface.
type MeasurableGroup map[string]Measurable

/* Measurable interface implementation */

func (h MeasurableGroup) GetPower() float64 {
	power := 0.0
	for _, consumer := range h {
		power += consumer.GetPower()
	}
	return power
}

func (h MeasurableGroup) GetEnergy() float64 {
	energy := 0.0
	for _, consumer := range h {
		energy += consumer.GetEnergy()
	}
	return energy
}

func (h MeasurableGroup) Reset() {
	for _, consumer := range h {
		consumer.Reset()
	}
}
