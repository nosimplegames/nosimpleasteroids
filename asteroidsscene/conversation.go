package asteroidsscene

import (
	"simple-games.com/asteroids/game"
	"simple-games.com/asteroids/scenes"
)

type Conversation struct {
	scenes.Act
	Player       *game.Player
	Conversation string
}

func (act *Conversation) Openning(scene *scenes.Scene) {
	act.Player.IsControllerEnabled = false
	conversation := game.TextDialogFactory{
		Conversation: act.Conversation,
		OnEnd:        act.Die,
	}.Create()

	act.AddChild(conversation)
}
