package nrgsim

import "time"

type simulation struct {
	consumers []simulable
	grid      gridConnection
}

// New creates a new electrical energy simulation.
func New() Simulation {
	return &simulation{}
}

/* Simulation interface implementation */

func (s *simulation) AddCustomConsumer(powerReadFunction PowerReadFunction) (Measurable, error) {
	c := &customConsumer{
		powerReadFunction: powerReadFunction,
	}
	s.consumers = append(s.consumers, c)
	return c, nil
}

func (s *simulation) AddCustomConsumerObject(nominalPower float64) (CustomConsumer, error) {
	c := &customConsumerObject{
		nominalPower: nominalPower,
	}
	c.powerReadFunction = c.getPower
	s.consumers = append(s.consumers, c)
	return c, nil
}

func (s *simulation) AddNoiseConsumer(minimumPower float64, maximumPower float64, interval time.Duration) (Measurable, error) {
	n := &noiseConsumer{
		minimumPower: minimumPower,
		maximumPower: maximumPower,
		interval:     interval,
		remaining:    interval,
	}
	s.consumers = append(s.consumers, n)
	return n, nil
}

func (s *simulation) AddNoiseConsumerObject(minimumPower float64, maximumPower float64, interval time.Duration) (NoiseConsumer, error) {
	n := &noiseConsumerObject{
		noiseConsumer{
			minimumPower: minimumPower,
			maximumPower: maximumPower,
			interval:     interval,
			remaining:    interval,
		},
	}
	s.consumers = append(s.consumers, n)
	return n, nil
}

func (s *simulation) AddControlledConsumer(nominalPower float64, controlFunction ControlFunction) (Measurable, error) {
	c := &controlledConsumer{
		nominalPower:    nominalPower,
		controlFunction: controlFunction,
	}
	s.consumers = append(s.consumers, c)
	return c, nil
}

func (s *simulation) AddControlledConsumerObject(nominalPower float64) (ControlledConsumer, error) {
	c := &controlledConsumerObject{
		nominalPower: nominalPower,
		isOn:         false,
	}
	c.controlFunction = c.IsOn
	s.consumers = append(s.consumers, c)
	return c, nil
}

func (s *simulation) AddAdjustableConsumer(nominalPower float64, adjustFunction AdjustFunction) (Measurable, error) {
	c := &adjustableProducerConsumer{
		isProducer:     false,
		nominalPower:   nominalPower,
		adjustFunction: adjustFunction,
	}
	s.consumers = append(s.consumers, c)
	return c, nil
}

func (s *simulation) AddAdjustableConsumerObject(nominalPower float64) (AdjustableConsumer, error) {
	c := &adjustableProducerConsumerObject{
		factor: 0.0,
	}
	c.isProducer = false
	c.nominalPower = nominalPower
	c.adjustFunction = c.GetFactor
	s.consumers = append(s.consumers, c)
	return c, nil

}

func (s *simulation) AddAdjustableProducer(nominalPower float64, adjustFunction AdjustFunction) (Measurable, error) {
	p := &adjustableProducerConsumer{
		isProducer:     true,
		nominalPower:   nominalPower,
		adjustFunction: adjustFunction,
	}
	s.consumers = append(s.consumers, p)
	return p, nil
}

func (s *simulation) AddAdjustableProducerObject(nominalPower float64) (AdjustableProducer, error) {
	p := &adjustableProducerConsumerObject{
		factor: 0.0,
	}
	p.isProducer = true
	p.nominalPower = nominalPower
	p.adjustFunction = p.GetFactor
	s.consumers = append(s.consumers, p)
	return p, nil
}

func (s *simulation) AddEnergyStorage(maxChargePower float64, maxDischargePower float64, capacity float64, control StorageControlFunction) (Storable, error) {
	e := &energyStorage{
		maxChargePower:    maxChargePower,
		maxDischargePower: maxDischargePower,
		capacity:          capacity,
		control:           control,
		stored:            capacity * 3 / 4,
	}
	s.consumers = append(s.consumers, e)
	return e, nil
}

func (s *simulation) AddEnergyStorageObject(maxChargePower float64, maxDischargePower float64, capacity float64) (EnergyStorage, error) {
	e := &energyStorageObject{
		energyStorage: energyStorage{
			maxChargePower:    maxChargePower,
			maxDischargePower: maxDischargePower,
			capacity:          capacity,
			stored:            capacity * 3 / 4,
		},
	}
	e.control = e.controlFunc
	s.consumers = append(s.consumers, e)
	return e, nil
}

func (s *simulation) GetGridConnection() Measurable {
	return &s.grid
}

func (s *simulation) CalculateStep(duration time.Duration) {
	powerSum := 0.0
	for _, c := range s.consumers {
		c.calculateStep(duration)
		powerSum += c.GetPower()
	}
	s.grid.setPower(powerSum, duration)
}

func (s *simulation) Reset() {
	for _, c := range s.consumers {
		c.Reset()
	}
	s.grid.Reset()
}
