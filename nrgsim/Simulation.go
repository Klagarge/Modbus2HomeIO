package nrgsim

import "time"

// Simulation is the interface that can be used to simulate the power and energy consumption of a simple household. It can be used to add consumers, producers and energy storage
// devices to the simulation. The simulation can then be run for a given duration.
type Simulation interface {
	// AddNoiseConsumer adds a customConsumer to the simulation that consumes a random amount of power between the minimum and maximum power values with new values at the given interval.
	// Note that the power values and the interval can not be changed after the customConsumer has been added to the simulation.
	AddNoiseConsumer(minimumPower float64, maximumPower float64, interval time.Duration) (Measurable, error)

	// AddNoiseConsumerObject adds a customConsumer to the simulation that consumes a random amount of power between the minimum and maximum power values with new values at the given
	// interval. The customConsumer can be controlled by using the NoiseConsumer interface.
	AddNoiseConsumerObject(minimumPower float64, maximumPower float64, interval time.Duration) (NoiseConsumer, error)

	// AddControlledConsumer adds a customConsumer to the simulation that can be controlled by using the ControlFunction interface. The customConsumer will consume the nominal
	// power if it is turned on and 0 W otherwise. The control function is called at every simulation step to check if the customConsumer is turned on or off, so it should be
	// implemented efficiently. The nominal power can not be changed after the customConsumer has been added to the simulation.
	AddControlledConsumer(nominalPower float64, control ControlFunction) (Measurable, error)

	// AddControlledConsumerObject adds a customConsumer to the simulation that can be controlled by using the ControlledConsumer interface. The customConsumer will consume the nominal
	// power if it is turned on and 0 W otherwise. The nominal power can be changed at any time. The customConsumer can be turned on or off at any time using the SetOn method of the
	// ControlledConsumer interface. The nominal power can be changed too using the SetNominalPower method of the ControlledConsumer interface.
	AddControlledConsumerObject(nominalPower float64) (ControlledConsumer, error)

	// AddAdjustableConsumer adds a consumer to the simulation that can be adjusted by using the AdjustFunction interface. The consumer will consume a fraction of the nominal power
	// returned by the adjust function. The adjust function is called at every simulation step to get the power level, so it should be implemented efficiently. The nominal power
	// can not be changed after the consumer has been added to the simulation.
	AddAdjustableConsumer(nominalPower float64, adjust AdjustFunction) (Measurable, error)

	// AddAdjustableConsumerObject adds an adjustable consumer to the simulation. The consumer can be controlled by using the AdjustableConsumer interface. The nominal power can be
	// changed at any time. The consumer's power level (0.0-1.0) can be changed at any time using the SetFactor method of the AdjustableConsumer interface.
	AddAdjustableConsumerObject(nominalPower float64) (AdjustableConsumer, error)

	// AddAdjustableProducer adds a producer to the simulation that can be adjusted by using the AdjustFunction interface. The producer will produce a fraction of the nominal power
	// returned by the adjust function. The adjust function is called at every simulation step to get the power level, so it should be implemented efficiently. The nominal power
	// can not be changed after the consumer has been added to the simulation.
	AddAdjustableProducer(nominalPower float64, adjust AdjustFunction) (Measurable, error)

	// AddAdjustableProducerObject adds an adjustable producer to the simulation. The producer can be controlled by using the AdjustableProducer interface. The nominal power can be
	// changed at any time. The producer's power level (0.0-1.0) can be changed at any time using the SetFactor method of the AdjustableProducer interface.
	AddAdjustableProducerObject(nominalPower float64) (AdjustableProducer, error)

	// AddCustomConsumer adds a customConsumer to the simulation. The power read function is called every time the simulation is calculated to check if the customConsumer is actually turned
	// on or off and what amount of power it is consuming.
	AddCustomConsumer(power PowerReadFunction) (Measurable, error)

	// AddCustomConsumerObject adds a customConsumer to the simulation. The customConsumer can be controlled by using the ConsumerObject interface.
	AddCustomConsumerObject(nominalPower float64) (CustomConsumer, error)

	// AddEnergyStorage adds an energy storage device to the simulation. The energy storage device can be controlled by using the StorageControlFunction callback. The energy storage
	// device will charge with the given power if the control function returns StorageCharge and discharge with the given power if the control function returns StorageDischarge. The
	// energy storage device will not charge or discharge if the control function returns StorageIdle. The energy storage device will not charge or discharge more than the given
	// maximum charge or discharge power. The energy storage device will not store more energy than the given capacity. The control function is called at every simulation step to
	//check if the energy storage device should charge or discharge and the current power limitation, so it should be implemented efficiently.
	AddEnergyStorage(maxChargePower float64, maxDischargePower float64, capacity float64, control StorageControlFunction) (Storable, error)

	// AddEnergyStorageObject adds an energy storage device to the simulation. The energy storage device can be controlled by using the EnergyStorage interface. The energy storage
	// device will charge with the given configuration. The energy storage device will not store more energy than the given capacity. The energy storage device can be controlled
	// at any time using the SetMode, the SetMaxChargePower and SetMaxDischargePower methods of the EnergyStorage interface.
	AddEnergyStorageObject(maxChargePower float64, maxDischargePower float64, capacity float64) (EnergyStorage, error)

	// GetGridConnection returns a Measurable object that represents the power consumption of the household from the POV of the grid. The power consumption of the household from
	// the grid is the sum of the power consumption of all consumers plus the (negative) sum of the power production of all producers.
	GetGridConnection() Measurable

	// CalculateStep calculates the power and energy consumption of the simulated household for the given duration.
	CalculateStep(duration time.Duration)

	// Reset resets the simulation to its initial state.
	Reset()
}
