package engine

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type IDrawable interface {
	IsVisible() bool
	Draw(target render.RenderTarget, combinedTransform math.Transform)
}
