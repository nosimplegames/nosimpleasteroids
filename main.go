package main

import (
	"simple-games.com/asteroids/asteroidsscene"
	"simple-games.com/asteroids/engine"
)

func main() {
	entities := engine.GetEntities()
	scene := asteroidsscene.AsteroidsSceneFactory{}.Create()
	entities.AddEntity(scene)

	asteroidsGame := engine.Game{}
	asteroidsGame.Run()
}
