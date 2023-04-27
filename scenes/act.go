package scenes

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/utils"
)

type Act struct {
	engine.Entity
}

func (act Act) MayBeSkipped() bool {
	return false
}

func (act *Act) OnDie(callback utils.Callback) {
	act.Entity.OnDie.AddCallback(callback)
}
