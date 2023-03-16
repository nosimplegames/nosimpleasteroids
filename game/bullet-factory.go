package game

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
)

type BulletFactory struct {
	Position math.Vector
	Rotation float64
}

func (factory BulletFactory) Create() *Bullet {
	bullet := &Bullet{}

	bullet.Position = factory.Position
	bullet.Origin = math.Vector{X: 5, Y: 2.5}
	bullet.Rotation = factory.Rotation
	bullet.LifeTime = 0.4
	bullet.Speed = 15

	engine.GetWorld().AddCollisinable(bullet)

	return bullet
}
