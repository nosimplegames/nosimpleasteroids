package asteroidsscene

import (
	"fmt"

	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/game"
	"simple-games.com/asteroids/scenes"
)

type BlastEnergy struct {
	scenes.Act

	Player *game.Player
}

func (act *BlastEnergy) Openning(scene *scenes.Scene) {
	blast := game.EnergyBlastFactory{}.Create()
	act.AddChild(blast)

	shield := act.Player.ActivateShield()

	shield.OnDie.AddCallback(func() {
		act.Player.Die()
		fmt.Println("Run destroy animation here")

		engine.GetTimer().AddTimeout(&engine.Timeout{
			Time:     3,
			Callback: act.Die,
		})
	})
}
