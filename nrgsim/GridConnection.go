package nrgsim

import "time"

type gridConnection struct {
	baseMeasurable
}

func (g *gridConnection) setPower(power float64, deltaTime time.Duration) {
	g.power = power
	g.energy += power * deltaTime.Seconds()
}
