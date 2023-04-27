package scenes

import (
	"simple-games.com/asteroids/game"
)

type Conversation struct {
	Act
	Player       *game.Player
	Conversation string
}

func (act *Conversation) Openning(_ *Scene) {
	act.Player.IsControllerEnabled = false
	conversation := game.TextDialogFactory{
		Conversation: "Where am I? Command Center is not responding and navigation system is not working at all... What is that?... Is it... an asteroid!?",
		OnEnd:        act.Die,
	}.Create()

	act.AddChild(conversation)
}
