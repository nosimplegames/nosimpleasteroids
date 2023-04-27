package scenes

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/utils"
)

type IAct interface {
	engine.IEntity

	MayBeSkipped() bool
	Openning(*Scene)
	OnDie(utils.Callback)
}
