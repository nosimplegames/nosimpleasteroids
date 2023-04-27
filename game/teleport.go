package game

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
	"simple-games.com/asteroids/math"
)

type Teleport struct {
	engine.Living

	Position      math.Vector
	Size          math.Vector
	CollisionMask string

	Target                 math.ITransformable
	TargetPosition         math.Vector
	TargetCollisionMask    string
	TeleportAfterCollision bool

	OnCollisionWithTarget events.Signal
}

func (teleport Teleport) GetSize() math.Vector {
	return teleport.Size
}

func (teleport Teleport) GetPosition() math.Vector {
	return teleport.Position
}

func (teleport Teleport) GetCollisionMask() string {
	return teleport.CollisionMask
}

func (teleport Teleport) CanCollide() bool {
	return true
}

func (teleport Teleport) CanCollideWith(collisionMask string) bool {
	return collisionMask == teleport.TargetCollisionMask
}

func (teleport Teleport) OnCollision(collisionMask string) {
	teleport.OnCollisionWithTarget.TFire(&teleport)
}

func (teleport Teleport) Teleport() {
	teleport.Target.SetPosition(teleport.TargetPosition)
}
