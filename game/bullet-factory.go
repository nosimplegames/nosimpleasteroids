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

	bullet.SetPosition(factory.Position)
	bullet.SetRotation(factory.Rotation)
	bullet.LifeTime = 1.5
	bullet.Speed = 40
	bullet.Texture = GetAssets().BulletTexture
	bullet.SetOriginCenter()

	engine.GetWorld().AddCollisinable(bullet)

	return bullet
}
