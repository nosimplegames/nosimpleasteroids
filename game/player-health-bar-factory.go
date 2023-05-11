package game

import (
	"simple-games.com/asteroids/assets"
	"simple-games.com/asteroids/ui/progressbar"
)

type PlayerHealthBarFactory struct {
}

func (factory PlayerHealthBarFactory) Create() *PlayerHealthBar {
	gameAssets := GetAssets()
	healthBar := &PlayerHealthBar{
		ProgressBar: *progressbar.Factory{
			ProgressTexture: gameAssets.ProgressTexture,
			BarTexture:      gameAssets.BarTexture,
		}.Create(),
	}
	healthBar.ProgressOffset = assets.HealthBarOffset

	healthBar.MaxValue = 100
	healthBar.Value = 100

	return healthBar
}
