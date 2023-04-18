package textdialog

import (
	"golang.org/x/image/font"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
	"simple-games.com/asteroids/ui"
	"simple-games.com/asteroids/utils"
	"simple-games.com/asteroids/utils/utilstext"
)

type Factory struct {
	Slice9            render.Slice9
	TextLinesCount    int
	LineWidth         float64
	Padding           ui.Padding
	FontFace          font.Face
	Text              string
	NextDialogTexture render.Texture
}

func (factory Factory) Create() *TextDialog {
	dialog := &TextDialog{}

	dialog.FontFace = factory.FontFace
	dialog.TextOffset = factory.GetTextPosition()
	targetSize := factory.GetTargetSize()
	dialog.BackgroundTexture = factory.Slice9.Compose(targetSize)
	dialog.Origin = targetSize.By(0.5)
	dialog.Dialogs = factory.GetDialogs()

	nextDialogEntity := factory.CreateNextDialogEntity()
	dialog.AddChild(nextDialogEntity)

	factory.ConnectEvents(dialog, nextDialogEntity)

	return dialog
}

func (factory Factory) GetTargetSize() math.Vector {
	lineHeight := utilstext.GetFontFaceHeight(factory.FontFace)

	return math.Vector{
		X: factory.LineWidth + factory.Padding.Right + factory.Padding.Left,
		Y: float64(lineHeight)*float64(factory.TextLinesCount) + factory.Padding.Top + factory.Padding.Bottom,
	}
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
	nextDialog.Position = factory.GetNextDialogPosition()

	return nextDialog
}

func (factory Factory) GetNextDialogPosition() math.Vector {
	size := factory.GetTargetSize()

	return size.ByVector(math.Vector{
		X: 0.93,
		Y: 0.75,
	})
}

func (factory Factory) GetDialogs() utils.List[Dialog] {
	return utils.List[Dialog]{
		Elements: DialogsFactory{
			TextLinesCount: factory.TextLinesCount,
			FontFace:       factory.FontFace,
			LineWidth:      factory.LineWidth,
			Text:           factory.Text,
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
