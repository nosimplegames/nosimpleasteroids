package asteroidsscene

import (
	"simple-games.com/asteroids/assets"
	"simple-games.com/asteroids/game"
	"simple-games.com/asteroids/huds"
	"simple-games.com/asteroids/scenes"
)

type AsteroidsSceneFactory struct {
}

func (factory AsteroidsSceneFactory) Create() *scenes.Scene {
	player := factory.createPlayer()

	scene := &scenes.Scene{
		Acts: scenes.Acts{
			Elements: scenes.ActsElements{
				// &UnknownReality{
				// 	Conversation: Conversation{
				// 		Player: player,
				// 	},
				// },
				&Asteroids{
					Player: player,
				},
				// &AfterAsteroids{
				// 	Conversation: Conversation{
				// 		Player: player,
				// 	},
				// },
				// &BlastEnergy{
				// 	Player: player,
				// },
				// &GameTitle{},
			},
		},
	}

	hud := huds.SpatialHudFactory{
		Player: player,
	}.Create()

	scene.AddChild(player)
	scene.AddChild(hud)

	return scene
}

func (scene *AsteroidsSceneFactory) createPlayer() *game.Player {
	player := game.PlayerFactory{
		IsControllerDisabled: true,
	}.Create()
	player.SetPosition(assets.GameSize.By(0.5))

	return player
}
