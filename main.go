package main

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/game/scenes"
)

func main() {
	entities := engine.GetEntities()
	scene := &scenes.AsteroidsScene{}
	scene.Init()
	entities.AddEntity(scene)
	// entities.AddEntity(game.PlayerFactory{}.Create())
	// entities.AddEntity(game.GetScore())

	// game.InitAsteroidsGenerator()

	asteroidsGame := engine.Game{}
	asteroidsGame.Run()
}
