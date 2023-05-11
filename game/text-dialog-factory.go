package game

import (
	"simple-games.com/asteroids/assets"
	"simple-games.com/asteroids/events"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/ui"
	"simple-games.com/asteroids/ui/textdialog"
	"simple-games.com/asteroids/utils/utilstext"
)

type TextDialogFactory struct {
	Conversation string
	OnEnd        func()
}

func (factory TextDialogFactory) Create() *textdialog.TextDialog {
	gameAssets := GetAssets()

	lineCount := 2
	lineHeight := 2.0
	lineSize := utilstext.GetFontFaceHeight(gameAssets.UIFontFace) * lineHeight
	dialogSize := math.Vector{
		X: assets.GameSize.X - assets.UIPadding*2,
		Y: lineSize*float64(lineCount) + assets.DialogPadding.Top + assets.DialogPadding.Bottom,
	}

	textDialog := textdialog.Factory{
		Slice9:            gameAssets.TextDialogSlice9,
		TextLinesCount:    lineCount,
		LineHeight:        lineHeight,
		Text:              factory.Conversation,
		FontFace:          gameAssets.UIFontFace,
		Size:              dialogSize,
		Padding:           assets.DialogPadding,
		NextDialogTexture: gameAssets.NextDialogTexture,
		Color:             assets.TextDialogColor,
	}.Create()
	textDialog.SetPosition(math.Vector{
		X: assets.GameSize.X * 0.5,
		Y: assets.GameSize.Y - assets.UIPadding - dialogSize.Y*0.5,
	})

	textDialog.AddEventListener(events.EventListener{
		EventType: ui.TextDialogFinishEvent,
		Callback:  func(event events.Event) { factory.OnEnd() },
	})

	return textDialog
}
