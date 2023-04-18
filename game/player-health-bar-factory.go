package game

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/ui/progressbar"
)

type PlayerHealthBarFactory struct {
}

func (factory PlayerHealthBarFactory) Create() *PlayerHealthBar {
	assets := GetAssets()
	healthBar := &PlayerHealthBar{
		ProgressBar: *progressbar.Factory{
			ProgressTexture: assets.ProgressTexture,
			BarTexture:      assets.BarTexture,
		}.Create(),
	}

	healthBar.Position = math.Vector{X: 20, Y: 10}
	healthBar.MaxValue = 3
	healthBar.Value = 3

	return healthBar
}
