package game

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
)

type Inert struct {
	engine.Sprite

	Speed float64
}

func (inert *Inert) Update() {
	movementVector := math.MovementVector{
		Speed:    inert.Speed,
		Rotation: inert.Rotation,
	}.Calculate()

	inert.Move(movementVector)
}
