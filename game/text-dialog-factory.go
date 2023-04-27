package game

import (
	"simple-games.com/asteroids/events"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/ui"
	"simple-games.com/asteroids/ui/textdialog"
)

type TextDialogFactory struct {
	Conversation string
	OnEnd        func()
}

func (factory TextDialogFactory) Create() *textdialog.TextDialog {
	assets := GetAssets()

	textDialog := textdialog.Factory{
		Slice9:         assets.TextDialogSlice9,
		TextLinesCount: 2,
		LineHeight:     2,
		Text:           factory.Conversation,
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
		Callback:  func(event events.Event) { factory.OnEnd() },
	})

	return textDialog
}
