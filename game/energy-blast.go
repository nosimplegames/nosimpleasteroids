package game

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
)

type EnergyBlast struct {
	Inert
	engine.Sprite
}

func (blast EnergyBlast) GetTransform() math.Transform {
	return blast.Inert.GetTransform()
}
