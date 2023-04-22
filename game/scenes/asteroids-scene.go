package scenes

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
	"simple-games.com/asteroids/game"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/ui"
	"simple-games.com/asteroids/ui/textdialog"
)

type AsteroidsScene struct {
	engine.Scene

	Player *game.Player
}

func (scene *AsteroidsScene) Init() {
	scene.createPlayer()
	scene.startInitialConversation()
}

func (scene *AsteroidsScene) createPlayer() {
	player := game.PlayerFactory{
		IsControllerDisabled: true,
	}.Create()

	scene.Player = player
	scene.AddChild(player)
}

func (scene *AsteroidsScene) startInitialConversation() {
	scene.startConversation("Where am I? Command Center is not responding and navigation system is not working at all... What is that?... Is it... an asteroids!?", scene.onInitialConversationEnds)
}

func (scene *AsteroidsScene) onInitialConversationEnds() {
	scene.Player.IsControllerEnabled = true
	game.InitAsteroidsGenerator()

	engine.GetTimer().AddTimeout(&engine.Timeout{
		Time: 10,
		Callback: func() {
			game.StopAsteroidsGenerator()
			scene.startConversationAfterAsteroids()
		},
	})
}

func (scene *AsteroidsScene) startConversationAfterAsteroids() {
	scene.Player.IsControllerEnabled = false
	scene.startConversation("Uff! I hope they were all. I haven't seen anything like that before. Mmm... scanner is showing something big ahead... damn that's definitely not a asteroid! Shields now!", func() {
		scene.createEnergyBlast()
	})
}

// Temporary method, must be replaced by textdialog.conversation
func (scene *AsteroidsScene) startConversation(conversation string, callback func()) {
	assets := game.GetAssets()

	textDialog := textdialog.Factory{
		Slice9:         assets.TextDialogSlice9,
		TextLinesCount: 2,
		LineHeight:     2,
		Text:           conversation,
		FontFace:       assets.UIFontFace,
		LineWidth:      300,
		Padding: ui.Padding{
			Top:    10,
			Bottom: 10,
			Left:   10,
			Right:  50,
		},
		NextDialogTexture: assets.NextDialogTexture,
	}.Create()
	textDialog.Position = math.Vector{
		X: 400,
		Y: 700,
	}

	textDialog.AddEventListener(events.EventListener{
		EventType: ui.TextDialogFinishEvent,
		Callback:  func(event events.Event) { callback() },
	})

	scene.AddChild(textDialog)
}

func (scene *AsteroidsScene) createEnergyBlast() {
	blast := game.EnergyBlastFactory{}.Create()
	scene.AddChild(blast)
	scene.Player.ActivateShield()
}

// func (scene *AsteroidsScene) createScore() {
// scene.AddChild(game.GetScore())
// }
