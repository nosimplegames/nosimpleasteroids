package game

import (
	"github.com/hajimehoshi/ebiten"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/particles"
	"simple-games.com/asteroids/render"
)

type AsteroidSize int

const (
	BigAsteroid AsteroidSize = iota
	NormalAsteroid
	SmallAsteroid
)

type OnExplodeFn = func(*Asteroid)

type Asteroid struct {
	engine.Entity
	Inert
	particles.Particle

	Size         AsteroidSize
	Texture      *ebiten.Image
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
		asteroid.LifeTime = 20
		asteroid.LifePoints = 3
		asteroid.SizeInPixels = math.Vector{X: 48, Y: 48}

	case NormalAsteroid:
		asteroid.Texture = assets.NormalAsteroidTexture
		asteroid.Speed = 1.75
		asteroid.LifeTime = 15
		asteroid.LifePoints = 2
		asteroid.SizeInPixels = math.Vector{X: 33, Y: 33}

	case SmallAsteroid:
		asteroid.Texture = assets.SmallAsteroidTexture
		asteroid.Speed = 3
		asteroid.LifeTime = 10
		asteroid.LifePoints = 1
		asteroid.SizeInPixels = math.Vector{X: 20, Y: 20}
	}

	asteroid.Origin = asteroid.SizeInPixels.By(0.5)
}

func (asteroid Asteroid) Draw(screen *ebiten.Image) {
	render.Sprite{
		Texture:   asteroid.Texture,
		Transform: asteroid.GetTransform(),
		Target:    screen,
	}.Render()
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
