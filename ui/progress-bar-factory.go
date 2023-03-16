package ui

import (
	"simple-games.com/asteroids/render"
)

type ProgressBarFactory struct {
	BarTexture      render.Texture
	ProgressTexture render.Texture
}

func (factory ProgressBarFactory) Create() *ProgressBar {
	progressBar := &ProgressBar{
		BarTexture:      factory.BarTexture,
		ProgressTexture: factory.ProgressTexture,
	}

	barTextureSize := render.GetTextureSize(factory.BarTexture)
	progressTextureSize := render.GetTextureSize(factory.ProgressTexture)

	progressBar.Origin = barTextureSize.By(0.5)
	progressBar.ProgressOffset = barTextureSize.Minus(progressTextureSize).By(0.5)

	return progressBar
}
