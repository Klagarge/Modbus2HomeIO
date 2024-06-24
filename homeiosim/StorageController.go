package homeiosim

import "Modbus2HomeIO/nrgsim"

type StorageController struct {
	Mode              nrgsim.StorageMode
	MaxChargePower    float64
	MaxDischargePower float64
}

func (s *StorageController) ControlFunction() (nrgsim.StorageMode, float64) {
	switch s.Mode {
	case nrgsim.StorageCharge:
		return nrgsim.StorageCharge, s.MaxChargePower
	case nrgsim.StorageDischarge:
		return nrgsim.StorageDischarge, s.MaxDischargePower
	default:
		return nrgsim.StorageIdle, 0
	}
}
