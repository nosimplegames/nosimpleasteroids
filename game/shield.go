package game

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
)

type Shield struct {
	engine.Sprite

	Size math.Vector
}

func (shield Shield) GetSize() math.Vector {
	return shield.Size
}

func (shield Shield) GetCollisionMask() string {
	return SpaceshipShieldCollisionMask
}

func (shield Shield) CanCollide() bool {
	return shield.IsAlive()
}

func (shield Shield) CanCollideWith(collisionMask string) bool {
	return collisionMask == EnergyBlastCollisionMask
}

func (shield *Shield) OnCollision(collisionMask string) {
	shield.Die()
}
