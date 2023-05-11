package progressbar

import (
	"simple-games.com/asteroids/render"
)

type Factory struct {
	BarTexture      render.Texture
	ProgressTexture render.Texture
}

func (factory Factory) Create() *ProgressBar {
	progressBar := &ProgressBar{
		BarTexture:      factory.BarTexture,
		ProgressTexture: factory.ProgressTexture,
	}

	barTextureSize := render.GetTextureSize(factory.BarTexture)
	progressTextureSize := render.GetTextureSize(factory.ProgressTexture)

	progressBar.Size = barTextureSize
	progressBar.SetOrigin(barTextureSize.By(0.5))
	progressBar.ProgressOffset = barTextureSize.Minus(progressTextureSize).By(0.5)

	return progressBar
}
