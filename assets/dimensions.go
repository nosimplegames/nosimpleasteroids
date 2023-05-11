package assets

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/ui"
)

var (
	GameSize = math.Vector{
		X: 1920,
		Y: 1080,
	}
	UIPadding     = 60.0
	DialogPadding = ui.Padding{
		Top:    18,
		Bottom: 18,
		Left:   18,
		Right:  150,
	}
	HealthBarOffset = math.Vector{
		X: 66,
		Y: 9,
	}
)
