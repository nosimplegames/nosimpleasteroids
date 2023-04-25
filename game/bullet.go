package game

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/particles"
)

type Bullet struct {
	Inert
	particles.Particle
}

func (bullet *Bullet) HandleInput() {
}

func (bullet *Bullet) Update() {
	bullet.Inert.Update()
	bullet.Live()
}

func (bullet Bullet) CanCollide() bool {
	return true
}

func (bullet Bullet) CanCollideWith(collisionMask string) bool {
	return collisionMask == AsteroidCollisionMask
}

func (bullet Bullet) GetCollisionMask() string {
	return BulletCollisionMask
}

func (bullet Bullet) GetSize() math.Vector {
	return math.Vector{X: 10, Y: 5}
}

func (bullet *Bullet) OnCollision(collisionMask string) {
	bullet.Die()
	GetScore().Increase(100)

	ExplosionFactory{
		Position:      bullet.Position,
		ExplosionSize: SmallExplosion,
	}.Create()
}

func (bullet Bullet) IsAlive() bool {
	return bullet.Particle.IsAlive()
}

func (bullet *Bullet) Die() {
	bullet.Particle.Die()
}
