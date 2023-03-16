package ui

import (
	"golang.org/x/image/font"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type TextDialog struct {
	engine.Entity
	math.Transformable

	BackgroundTexture render.Texture
	FontFace          *font.Face
	Dialogs           []string
}

func (dialog TextDialog) Draw(target render.RenderTarget) {
	render.Sprite{
		Texture:   dialog.BackgroundTexture,
		Transform: dialog.GetTransform(),
		Target:    target,
	}.Render()
}
