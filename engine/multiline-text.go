package engine

import (
	"golang.org/x/image/font"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type MultilineText struct {
	Entity
	Colorable

	Lines      []string
	LineHeight float64
	FontFace   font.Face
}

func (text MultilineText) Draw(target render.RenderTarget, transform math.Transform) {
	render.MultilineText{
		Lines:      text.Lines,
		Target:     target,
		Position:   text.GetPosition(),
		ColorM:     text.ColorM,
		LineHeight: text.LineHeight,
		FontFace:   text.FontFace,
	}.Render()
}
