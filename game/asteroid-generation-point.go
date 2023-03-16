package game

import (
	"math/rand"

	"simple-games.com/asteroids/math"
)

type AsteroidGenerationPoint struct {
	Position    math.Vector
	MinRotation float64
	MaxRotation float64
}

func (point AsteroidGenerationPoint) GetRandomRotation() float64 {
	return point.MinRotation + rand.Float64()*(point.MaxRotation-point.MinRotation)
}
