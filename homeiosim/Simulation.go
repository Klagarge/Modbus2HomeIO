package homeiosim

import (
	"Modbus2HomeIO/homeio"
	"Modbus2HomeIO/nrgsim"
)

// Simulation is the interface to a Home I/O energy simulation. Using this interface, the user can interact with a Home I/O energy simulation by getting the grid connection, the
// PV producer, the storage, the heating consumers, and the total heating consumption. The user can also step the simulation forward in time, and reset the simulation to its
// initial state. Note that the simulation can run accelerated, meaning that the user can step the simulation forward in time by a duration that is different from the real time,
// depending on the time in the Home I/O game.
type Simulation interface {
	// GetGridConnection returns the grid connection of the simulated home.
	GetGridConnection() nrgsim.Measurable

	// GetPVProducer returns the PV producer of the simulated home.
	GetPVProducer() nrgsim.Measurable

	// GetStorageController returns the storage controller of the simulated home.
	GetStorageController() *StorageController

	// GetStorage returns the storage of the simulated home.
	GetStorage() nrgsim.Storable

	// GetHeatingConsumer returns the heating consumer of the given room.
	GetHeatingConsumer(room homeio.Room) (nrgsim.Measurable, error)

	// GetHeatingConsumersTotal returns the total heating consumption of the simulated home.
	GetHeatingConsumersTotal() nrgsim.Measurable

	// Step steps the simulation forward in time by the given duration.
	Step()

	// Reset resets the simulation to its initial state.
	Reset()
}
