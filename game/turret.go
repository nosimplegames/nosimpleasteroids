package game

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
)

type Turret struct {
	engine.Entity
	Weapon
}

func (turret *Turret) Shoot(position math.Vector, rotation float64) {
	if !turret.CanShoot() {
		return
	}

	turret.TimeSinceLastShoot = 0

	bullet := BulletFactory{
		Position: position,
		Rotation: rotation,
	}.Create()
	engine.GetEntities().AddEntity(bullet)
}
