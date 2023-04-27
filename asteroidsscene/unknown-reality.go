package asteroidsscene

import (
	"simple-games.com/asteroids/scenes"
)

type UnknownReality struct {
	Conversation
}

func (act *UnknownReality) Openning(scene *scenes.Scene) {
	act.Conversation.Conversation = "Where am I? Command Center is not responding and navigation system is not working at all... What is that?... Is it... an asteroid!?"
	act.Conversation.Openning(scene)
}
