package asteroidsscene

import (
	"simple-games.com/asteroids/game"
	"simple-games.com/asteroids/scenes"
)

type AsteroidsSceneFactory struct {
}

func (factory AsteroidsSceneFactory) Create() *scenes.Scene {
	player := factory.createPlayer()

	scene := &scenes.Scene{
		Acts: scenes.Acts{
			Elements: scenes.ActsElements{
				&UnknownReality{
					Conversation: Conversation{
						Player: player,
					},
				},
				&Asteroids{
					Player: player,
				},
				&AfterAsteroids{
					Conversation: Conversation{
						Player: player,
					},
				},
				&BlastEnergy{
					Player: player,
				},
				&GameTitle{},
			},
		},
	}

	scene.AddChild(player)

	return scene
}

func (scene *AsteroidsSceneFactory) createPlayer() *game.Player {
	player := game.PlayerFactory{
		IsControllerDisabled: true,
	}.Create()

	return player
}
