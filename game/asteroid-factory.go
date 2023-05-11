package game

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
)

type AsteroidFactory struct {
	Size     AsteroidDimensions
	Position math.Vector
	Rotation float64
}

func (factory AsteroidFactory) Create() *Asteroid {
	asteroid := &Asteroid{}

	asteroid.SetSize(factory.Size)
	asteroid.SetRotation(factory.Rotation)
	asteroid.SetPosition(factory.Position)
	asteroid.OnExplode = OnAsteroidExplode
	asteroid.Type = AsteroidEntityType

	engine.GetWorld().AddCollisinable(asteroid)

	return asteroid
}
