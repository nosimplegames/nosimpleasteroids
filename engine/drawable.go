package engine

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type Drawable struct {
	IsHidden bool
}

func (drawable Drawable) Draw(target render.RenderTarget, combinedTransform math.Transform) {
}

func (drawable *Drawable) IsVisible() bool {
	return !drawable.IsHidden
}

func (drawable *Drawable) Hide() {
	drawable.IsHidden = true
}

func (drawable *Drawable) Show() {
	drawable.IsHidden = false
}
