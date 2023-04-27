package main

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/game/asteroidsscene"
)

func main() {
	entities := engine.GetEntities()
	scene := asteroidsscene.AsteroidsSceneFactory{}.Create()
	entities.AddEntity(scene)
	// entities.AddEntity(game.PlayerFactory{}.Create())
	// entities.AddEntity(game.GetScore())

	// game.InitAsteroidsGenerator()

	asteroidsGame := engine.Game{}
	asteroidsGame.Run()
}
