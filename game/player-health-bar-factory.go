package game

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/ui"
)

type PlayerHealthBarFactory struct {
}

func (factory PlayerHealthBarFactory) Create() *PlayerHealthBar {
	assets := GetAssets()
	healthBar := &PlayerHealthBar{
		ProgressBar: *ui.ProgressBarFactory{
			ProgressTexture: assets.ProgressTexture,
			BarTexture:      assets.BarTexture,
		}.Create(),
	}

	healthBar.Position = math.Vector{X: 20, Y: 10}
	healthBar.MaxValue = 3
	healthBar.Value = 3

	return healthBar
}
