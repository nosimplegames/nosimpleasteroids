package game

import "simple-games.com/asteroids/engine"

type AsteroidExplodeHandler struct {
	ExplodingAsteroid *Asteroid
}

func (handler AsteroidExplodeHandler) Handle() {
	handler.CreateSmallerAsteroids()
	handler.CreateExplosion()
}

func (handler AsteroidExplodeHandler) CreateSmallerAsteroids() {
	mayCreateSmallerAsteroids := handler.ExplodingAsteroid.Size != SmallAsteroid

	if !mayCreateSmallerAsteroids {
		return
	}

	rotation := handler.ExplodingAsteroid.Rotation + 0.785398

	for i := 0; i < 2; i++ {
		asteroid := AsteroidFactory{
			Rotation: rotation + float64(180*i),
			Position: handler.ExplodingAsteroid.Position,
			Size:     handler.ExplodingAsteroid.GetSmallerSize(),
		}.Create()

		engine.GetEntities().AddEntity(asteroid)
	}
}

func (handler AsteroidExplodeHandler) CreateExplosion() {
	mayCreateExplosion := handler.ExplodingAsteroid.Size == SmallAsteroid

	if !mayCreateExplosion {
		return
	}

	ExplosionFactory{
		Position:      handler.ExplodingAsteroid.Position,
		ExplosionSize: BigExplosion,
	}.Create()
}

func OnAsteroidExplode(explodingAsteroid *Asteroid) {
	AsteroidExplodeHandler{
		ExplodingAsteroid: explodingAsteroid,
	}.Handle()
}
