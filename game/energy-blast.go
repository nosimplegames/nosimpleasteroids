package game

import (
	"simple-games.com/asteroids/math"
)

type EnergyBlast struct {
	Inert
	Size math.Vector
}

func (blast EnergyBlast) GetSize() math.Vector {
	return blast.Size
}

func (blast EnergyBlast) GetCollisionMask() string {
	return EnergyBlastCollisionMask
}

func (blast EnergyBlast) CanCollide() bool {
	return true
}

func (blast EnergyBlast) CanCollideWith(collisionMask string) bool {
	return collisionMask == SpaceshipCollisionMask ||
		collisionMask == SpaceshipShieldCollisionMask
}

func (blast EnergyBlast) OnCollision(collisionMask string) {
}

func (blast EnergyBlast) IsAlive() bool {
	return true
}
