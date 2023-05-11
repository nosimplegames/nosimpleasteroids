package progressbar

import (
	"image"

	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type ProgressBar struct {
	engine.Entity

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

	progressTransfrom := bar.GetProgressTransform()
	progressTransfrom.Concat(transform)

	render.Sprite{
		Texture:   bar.GetProgressTexture(),
		Target:    target,
		Transform: progressTransfrom,
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
	transform := math.Transform{}
	transform.Translate(bar.ProgressOffset.X, bar.ProgressOffset.Y)

	return transform
}

func (bar *ProgressBar) SetValue(value float64) {
	bar.Value = value
}

func (bar *ProgressBar) SetMaxValue(value float64) {
	bar.Value = value
}
