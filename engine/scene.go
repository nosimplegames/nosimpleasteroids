package engine

import "simple-games.com/asteroids/math"

type Scene struct {
	Entity
}

func (scene Scene) GetTransform() math.Transform {
	return math.Transform{}
}
