package engine

import (
	"simple-games.com/asteroids/math"
)

type Colorable struct {
	ColorM math.ColorM
}

func (colorable *Colorable) SetAlpha(alpha float64) {
	colorable.ColorM.SetElement(3, 3, alpha)
}
