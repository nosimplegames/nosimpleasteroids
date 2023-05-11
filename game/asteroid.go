package game

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/particles"
)

type AsteroidDimensions int

const (
	BigAsteroid AsteroidDimensions = iota
	NormalAsteroid
	SmallAsteroid
)

type OnExplodeFn = func(*Asteroid)

const AsteroidEntityType = "asteroid"

type Asteroid struct {
	Inert
	particles.Particle

	Dimensions AsteroidDimensions
	LifePoints int

	OnExplode OnExplodeFn
}

func (asteroid Asteroid) GetSmallerSize() AsteroidDimensions {
	if asteroid.Dimensions == NormalAsteroid {
		return SmallAsteroid
	}

	return NormalAsteroid
}

func (asteroid *Asteroid) Update() {
	asteroid.Live()
	asteroid.Inert.Update()
}

func (asteroid *Asteroid) SetSize(size AsteroidDimensions) {
	asteroid.Dimensions = size
	assets := GetAssets()

	switch size {
	case BigAsteroid:
		asteroid.Texture = assets.BigAsteroidTexture
		asteroid.Speed = 1.5
		asteroid.LifeTime = 30
		asteroid.LifePoints = 9
		asteroid.Size = math.Vector{X: 72, Y: 72}

	case NormalAsteroid:
		asteroid.Texture = assets.NormalAsteroidTexture
		asteroid.Speed = 1.75
		asteroid.LifeTime = 24
		asteroid.LifePoints = 6
		asteroid.Size = math.Vector{X: 51, Y: 51}

	case SmallAsteroid:
		asteroid.Texture = assets.SmallAsteroidTexture
		asteroid.Speed = 3
		asteroid.LifeTime = 15
		asteroid.LifePoints = 3
		asteroid.Size = math.Vector{X: 36, Y: 36}
	}

	asteroid.SetOriginCenter()
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
