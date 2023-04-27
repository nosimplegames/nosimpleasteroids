package asteroidsscene

import (
	"simple-games.com/asteroids/scenes"
)

type AfterAsteroids struct {
	Conversation
}

func (act *AfterAsteroids) Openning(scene *scenes.Scene) {
	act.Conversation.Conversation = "Uff! I hope they were all. I haven't seen anything like that before. Mmm... scanner is showing something big ahead... damn that's definitely not a asteroid! Shields now!"
	act.Conversation.Openning(scene)
}
