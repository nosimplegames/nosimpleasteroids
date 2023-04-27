package asteroidsscene

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/game"
	"simple-games.com/asteroids/scenes"
)

type Asteroids struct {
	scenes.Act

	Player *game.Player
}

func (act *Asteroids) Openning(_ *scenes.Scene) {
	act.Player.IsControllerEnabled = true
	game.InitAsteroidsGenerator()

	engine.GetTimer().AddTimeout(&engine.Timeout{
		Time: 10,
		Callback: func() {
			game.StopAsteroidsGenerator()
			act.Die()
		},
	})
}
