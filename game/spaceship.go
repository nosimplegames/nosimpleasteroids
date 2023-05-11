package game

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
)

const SpaceshipEntityType = "spaceship"

type Spaceship struct {
	engine.Sprite
	HasLifePoints

	Propulsor    Engine
	RightRotator Engine
	LeftRotator  Engine

	Turret          IWeapon
	ShieldGenerator ShieldGenerator

	OnLifePointsChanged events.SignalTCallback
	IsRespawning        bool
}

func (spaceship Spaceship) IsAlive() bool {
	return spaceship.lifePoints > 0 && !spaceship.IsDead
}

func (spaceship Spaceship) CanCollide() bool {
	return !spaceship.IsRespawning && !spaceship.ShieldGenerator.IsShieldActivated
}

func (spaceship Spaceship) CanCollideWith(collisionMask string) bool {
	return collisionMask == AsteroidCollisionMask || collisionMask == SpaceshipTeleportCollisionMask
}

func (spaceship Spaceship) GetCollisionMask() string {
	return SpaceshipCollisionMask
}

func (spaceship *Spaceship) OnCollision(collisionMask string) {
	isCollidingWithTeleport := collisionMask == SpaceshipTeleportCollisionMask

	if isCollidingWithTeleport {
		return
	}

	spaceship.DecreaseLifePoints(33)

	gotDestroyed := spaceship.lifePoints == 0

	if gotDestroyed {
		spaceship.Die()
	}
}

func (spaceship *Spaceship) ActivateShield() *Shield {
	shield := spaceship.ShieldGenerator.ActivateShield()
	spaceship.AddChild(shield)

	return shield
}
