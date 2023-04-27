package game

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/particles"
)

type AsteroidSize int

const (
	BigAsteroid AsteroidSize = iota
	NormalAsteroid
	SmallAsteroid
)

type OnExplodeFn = func(*Asteroid)

const AsteroidEntityType = "asteroid"

type Asteroid struct {
	Inert
	particles.Particle

	Size         AsteroidSize
	SizeInPixels math.Vector
	LifePoints   int

	OnExplode OnExplodeFn
}

func (asteroid Asteroid) GetSmallerSize() AsteroidSize {
	if asteroid.Size == NormalAsteroid {
		return SmallAsteroid
	}

	return NormalAsteroid
}

func (asteroid *Asteroid) Update() {
	asteroid.Live()
	asteroid.Inert.Update()
}

func (asteroid *Asteroid) SetSize(size AsteroidSize) {
	asteroid.Size = size
	assets := GetAssets()

	switch size {
	case BigAsteroid:
		asteroid.Texture = assets.BigAsteroidTexture
		asteroid.Speed = 1.5
		asteroid.LifeTime = 10
		asteroid.LifePoints = 3
		asteroid.SizeInPixels = math.Vector{X: 48, Y: 48}

	case NormalAsteroid:
		asteroid.Texture = assets.NormalAsteroidTexture
		asteroid.Speed = 1.75
		asteroid.LifeTime = 8
		asteroid.LifePoints = 2
		asteroid.SizeInPixels = math.Vector{X: 33, Y: 33}

	case SmallAsteroid:
		asteroid.Texture = assets.SmallAsteroidTexture
		asteroid.Speed = 3
		asteroid.LifeTime = 5
		asteroid.LifePoints = 1
		asteroid.SizeInPixels = math.Vector{X: 20, Y: 20}
	}

	asteroid.Origin = asteroid.SizeInPixels.By(0.5)
}

func (asteroid Asteroid) CanCollide() bool {
	return true
}

func (asteroid Asteroid) CanCollideWith(collisionMask string) bool {
	return collisionMask == BulletCollisionMask ||
		collisionMask == SpaceshipCollisionMask
}

func (asteroid *Asteroid) GetCollisionMask() string {
	return AsteroidCollisionMask
}

func (asteroid *Asteroid) GetSize() math.Vector {
	return asteroid.SizeInPixels
}

func (asteroid *Asteroid) OnCollision(collisionMask string) {
	asteroid.LifePoints -= 1

	mayExplode := asteroid.LifePoints == 0

	if mayExplode {
		asteroid.OnExplode(asteroid)
	}
}

func (asteroid Asteroid) IsAlive() bool {
	return asteroid.Particle.IsAlive() && asteroid.LifePoints > 0
}

func (asteroid *Asteroid) Die() {
	asteroid.Particle.Die()
}
