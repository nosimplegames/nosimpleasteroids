package main

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/game"
	"simple-games.com/asteroids/ui"
)

func main() {
	entities := engine.GetEntities()
	// entities.AddEntity(game.PlayerFactory{}.Create())
	// entities.AddEntity(game.GetScore())

	// game.InitAsteroidsGenerator()

	textureDialog := ui.TextDialogFactory{
		Slice9: game.GetAssets().TextDialogSlice9,
	}.Create()

	entities.AddEntity(textureDialog)

	asteroidsGame := engine.Game{}
	asteroidsGame.Run()
}
