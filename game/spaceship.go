package game

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
	"simple-games.com/asteroids/math"
)

type Spaceship struct {
	events.EventTarget
	engine.Sprite

	Propulsor    Engine
	RightRotator Engine
	LeftRotator  Engine

	Turret          IWeapon
	ShieldGenerator ShieldGenerator

	LifePoints   int
	IsRespawning bool
}

func (spaceship Spaceship) IsAlive() bool {
	return spaceship.LifePoints > 0
}

func (spaceship Spaceship) GetSize() math.Vector {
	return math.Vector{
		X: 12,
		Y: 12,
	}
}

func (spaceship Spaceship) CanCollide() bool {
	return !spaceship.IsRespawning && !spaceship.ShieldGenerator.IsShieldActivated
}

func (spaceship Spaceship) CanCollideWith(collisionMask string) bool {
	return collisionMask == AsteroidCollisionMask
}

func (spaceship Spaceship) GetCollisionMask() string {
	return SpaceshipCollisionMask
}

func (spaceship *Spaceship) OnCollision(collisionMask string) {
	spaceship.LifePoints -= 1
	spaceship.DispatchEvent(events.Event{
		Type: LifePointsChanged,
		Data: spaceship.LifePoints,
	})

	gotDestroyed := spaceship.LifePoints == 0

	if gotDestroyed {
		spaceship.Die()
	}
}

func (spaceship *Spaceship) ActivateShield() *Shield {
	shield := spaceship.ShieldGenerator.ActivateShield()
	spaceship.AddChild(shield)

	return shield
}
