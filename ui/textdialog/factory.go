package textdialog

import (
	"image/color"

	"golang.org/x/image/font"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
	"simple-games.com/asteroids/ui"
	"simple-games.com/asteroids/utils"
)

type Factory struct {
	Slice9            render.Slice9
	TextLinesCount    int
	Size              math.Vector
	Padding           ui.Padding
	FontFace          font.Face
	LineHeight        float64
	Text              string
	NextDialogTexture render.Texture
	Color             color.Color
}

func (factory Factory) Create() *TextDialog {
	dialog := &TextDialog{}

	dialog.FontFace = factory.FontFace
	dialog.TextOffset = factory.GetTextPosition()
	dialog.BackgroundTexture = factory.Slice9.Compose(factory.Size)
	dialog.SetOrigin(factory.Size.By(0.5))
	dialog.Dialogs = factory.GetDialogs()
	dialog.LineHeight = factory.LineHeight
	dialog.SetColor(factory.Color)

	nextDialogEntity := factory.CreateNextDialogEntity()
	dialog.AddChild(nextDialogEntity)

	factory.ConnectEvents(dialog, nextDialogEntity)

	return dialog
}

func (factory Factory) GetTextPosition() math.Vector {
	return math.Vector{
		X: (factory.Padding.Left - factory.Padding.Right) / 2.0,
		Y: (factory.Padding.Top - factory.Padding.Bottom) / 2.0,
	}
}

func (factory Factory) CreateNextDialogEntity() engine.IEntity {
	nextDialog := &engine.Sprite{
		Texture: factory.NextDialogTexture,
	}
	nextDialog.SetOriginCenter()
	nextDialog.SetPosition(factory.GetNextDialogPosition())

	return nextDialog
}

func (factory Factory) GetNextDialogPosition() math.Vector {
	return factory.Size.ByVector(math.Vector{
		X: 0.93,
		Y: 0.75,
	})
}

func (factory Factory) GetDialogs() utils.List[Dialog] {
	lineWidth := factory.Size.X - factory.Padding.Left - factory.Padding.Right

	return utils.List[Dialog]{
		Elements: DialogsFactory{
			TextLinesCount: factory.TextLinesCount,
			FontFace:       factory.FontFace,
			LineWidth:      lineWidth,
			Text:           factory.Text,
			LineHeight:     factory.LineHeight,
		}.Create(),
	}
}

func (factory Factory) ConnectEvents(dialog *TextDialog, nextDialogEntity engine.IEntity) {
	dialog.AddEventListener(events.EventListener{
		EventType: ui.TextDialogLastElementEvent,
		Callback: func(_ events.Event) {
			nextDialogEntity.Die()
		},
	})
}
