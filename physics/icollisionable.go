package physics

import "simple-games.com/asteroids/math"

type ICollisionable interface {
	GetSize() math.Vector
	GetPosition() math.Vector
	OnCollision(collisionMask string)
	CanCollide() bool
	CanCollideWith(collisionMask string) bool
	GetCollisionMask() string
	IsAlive() bool
}
