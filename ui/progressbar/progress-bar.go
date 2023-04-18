package progressbar

import (
	"image"

	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type ProgressBar struct {
	engine.Entity
	math.Transformable

	BarTexture      render.Texture
	ProgressTexture render.Texture
	ProgressOffset  math.Vector

	MaxValue float64
	Value    float64
}

func (bar ProgressBar) Draw(target render.RenderTarget, transform math.Transform) {
	render.Sprite{
		Texture:   bar.BarTexture,
		Target:    target,
		Transform: transform,
	}.Render()

	render.Sprite{
		Texture:   bar.GetProgressTexture(),
		Target:    target,
		Transform: bar.GetProgressTransform(),
	}.Render()
}

func (bar ProgressBar) GetProgressTexture() render.Texture {
	textureSize := render.GetTextureSize(bar.ProgressTexture)
	rect := image.Rect(
		0,
		0,
		int(textureSize.X*bar.Value/bar.MaxValue),
		int(textureSize.Y),
	)

	return bar.ProgressTexture.SubImage(rect).(render.Texture)
}

func (bar ProgressBar) GetProgressTransform() math.Transform {
	transform := bar.GetTransform()
	transform.Translate(bar.ProgressOffset.X, bar.ProgressOffset.Y)

	return transform
}
