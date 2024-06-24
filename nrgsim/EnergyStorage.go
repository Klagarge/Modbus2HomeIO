package nrgsim

import (
	"math"
	"time"
)

type (
	StorageMode int

	StorageControlFunction func() (StorageMode, float64)

	energyStorage struct {
		baseMeasurable
		control StorageControlFunction

		maxChargePower    float64 // W
		maxDischargePower float64 // W
		capacity          float64 // J
		stored            float64 // J
	}

	EnergyStorage interface {
		Storable

		// GetMaxChargePower returns the maximum power that can be used to charge the energy storage in watts.
		GetMaxChargePower() float64

		// SetMaxChargePower allows to change the maximum power (watts) that can be used to charge the energy storage at any time.
		SetMaxChargePower(power float64)

		// GetMaxDischargePower returns the maximum power that can be used to discharge the energy storage in watts.
		GetMaxDischargePower() float64

		// SetMaxDischargePower allows to change the maximum power (watts) that can be used to discharge the energy storage at any time.
		SetMaxDischargePower(power float64)

		// GetMode returns the current mode of the energy storage.
		GetMode() StorageMode

		// SetMode allows to change the mode of the energy storage at any time. The energy storage will charge if the mode is StorageCharge, discharge if the mode is
		// StorageDischarge and do nothing if the mode is StorageIdle.
		SetMode(mode StorageMode)
	}

	energyStorageObject struct {
		energyStorage
		mode StorageMode
	}
)

const (
	StorageCharge StorageMode = iota
	StorageDischarge
	StorageIdle
)

/* simulable interface implementation for energyStorage */

func (e *energyStorage) calculateStep(delta time.Duration) {
	// Get control command and power from the control function.
	cmd, pwr := e.control()

	switch cmd {
	case StorageCharge:
		// Charge the energy storage if there is capacity left.
		if e.stored < e.capacity {
			// Calculate the power that can be used to charge the energy storage using the device's limits.
			e.power = math.Min(pwr, e.maxChargePower)

			// Update the energy and stored energy of the energy storage.
			e.energy += e.power * delta.Seconds()
			e.stored += e.maxChargePower * delta.Seconds()

			// Limit the stored energy to the capacity.
			if e.stored > e.capacity {
				e.stored = e.capacity
			}
		} else {
			// The energy storage is full, so it can not be charged.
			e.power = 0
		}

	case StorageDischarge:
		// Discharge the energy storage if there is energy stored.
		if e.stored > 0 {

			// Calculate the power that can be used to discharge the energy storage using the device's limits.
			e.power = -1 * math.Min(pwr, e.maxDischargePower)

			// Update the energy and stored energy of the energy storage.
			e.energy += e.power * delta.Seconds()
			e.stored += e.maxDischargePower * delta.Seconds()

			// Limit the stored energy to 0.
			if e.stored < 0 {
				e.stored = 0
			}
		} else {
			// The energy storage is empty, so it can not be discharged.
			e.power = 0
		}

	case StorageIdle:
		// The energy storage is idle, so it does not charge or discharge.
		e.power = 0
	}
}

func (e *energyStorage) Reset() {
	e.baseMeasurable.Reset()
	e.stored = e.capacity * 3 / 4
}

/* Storable interface implementation for energyStorage */

func (e *energyStorage) GetStateOfCharge() float64 {
	return e.stored / e.capacity
}

func (e *energyStorage) GetStoredEnergy() float64 {
	return e.stored
}

func (e *energyStorage) GetCapacity() float64 {
	return e.capacity
}

/* EnergyStorage interface implementation for energyStorageObject */

func (e *energyStorageObject) GetMaxChargePower() float64 {
	return e.maxChargePower
}

func (e *energyStorageObject) SetMaxChargePower(power float64) {
	e.maxChargePower = power
}

func (e *energyStorageObject) GetMaxDischargePower() float64 {
	return e.maxDischargePower
}

func (e *energyStorageObject) SetMaxDischargePower(power float64) {
	e.maxDischargePower = power
}

func (e *energyStorageObject) GetMode() StorageMode {
	return e.mode
}

func (e *energyStorageObject) SetMode(mode StorageMode) {
	e.mode = mode
}

func (e *energyStorageObject) controlFunc() (StorageMode, float64) {
	switch e.mode {
	case StorageCharge:
		return StorageCharge, e.maxChargePower
	case StorageDischarge:
		return StorageDischarge, e.maxDischargePower
	default:
		return StorageIdle, 0
	}
}
