package nrgsim

import "time"

// simulable is an interface that defines the methods that need to be implemented by objects that can be simulated. The calculateStep method should calculate the next step of the
// simulation based on the duration of the step, and the object should implement the Measurable interface.
type simulable interface {
	Measurable

	// calculateStep calculates the next state of the simulated entity based on the duration of the step.
	calculateStep(duration time.Duration)
}
