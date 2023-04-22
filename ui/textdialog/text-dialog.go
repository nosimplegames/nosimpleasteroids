package textdialog

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"golang.org/x/image/font"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
	"simple-games.com/asteroids/ui"
	"simple-games.com/asteroids/utils"
)

type TextDialog struct {
	engine.Entity
	math.Transformable
	events.EventTarget

	BackgroundTexture render.Texture
	FontFace          font.Face
	LineHeight        float64
	Dialogs           utils.List[Dialog]
	TextOffset        math.Vector
}

func (dialog *TextDialog) HandleInput() {
	mustMoveDialog := inpututil.IsKeyJustReleased(ebiten.KeyJ)

	if mustMoveDialog {
		dialog.Dialogs.Next()

		if dialog.Dialogs.IsLastElement() {
			dialog.DispatchLastElementEvent()
		} else if dialog.Dialogs.HasReachEnd() {
			dialog.DispatchFinishEvent()
			dialog.Die()
		}
	}
}

func (dialog TextDialog) Draw(target render.RenderTarget, transform math.Transform) {
	render.Sprite{
		Texture:   dialog.BackgroundTexture,
		Transform: transform,
		Target:    target,
	}.Render()

	render.MultilineText{
		Lines:      dialog.Dialogs.GetCurrent().Lines,
		FontFace:   dialog.FontFace,
		LineHeight: dialog.LineHeight,
		Target:     target,
		Position:   dialog.Position.Add(dialog.TextOffset),
	}.Render()
}

func (dialog TextDialog) DispatchLastElementEvent() {
	event := events.Event{
		Type: ui.TextDialogLastElementEvent,
	}
	dialog.DispatchEvent(event)
}

func (dialog TextDialog) DispatchFinishEvent() {
	event := events.Event{
		Type: ui.TextDialogFinishEvent,
	}
	dialog.DispatchEvent(event)
}
