package ui

import (
	"golang.org/x/image/font"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type TextDialogFactory struct {
	Slice9     render.Slice9
	TextLines  int
	LineWidth  float64
	LineHeight float64
	Padding    Padding
	FontFace   font.Face
	Text       string
}

func (factory TextDialogFactory) Create() *TextDialog {
	dialog := &TextDialog{}

	lineHeight := float64(factory.FontFace.Metrics().Height) * factory.LineHeight
	targetSize := math.Vector{
		X: factory.LineWidth + factory.Padding.Right + factory.Padding.Left,
		Y: float64(lineHeight)*lineHeight + factory.Padding.Top + factory.Padding.Bottom,
	}
	dialog.BackgroundTexture = factory.Slice9.Compose(targetSize)
	textLines, _ := render.TextLinesCalculator{
		Text:       factory.Text,
		LineHeight: lineHeight,
		LineWidth:  factory.LineWidth,
		FontFace:   factory.FontFace,
	}.Calculate()
	dialog.Dialogs = textLines
	// dialog.BackgroundTexture, _ = ebiten.NewImage(
	// 	int(size),
	// 	int(size),
	// 	ebiten.FilterDefault,
	// )
	// render.TextureRepeater{
	// 	Target:  dialog.BackgroundTexture,
	// 	Texture: assets.LoadTexture("./res/text-dialog.png"),
	// 	Area: math.Box{
	// 		Position: math.Vector{
	// 			X: 0,
	// 			Y: 0,
	// 		},
	// 		Size: math.Vector{size, size},
	// 	},
	// }.Repeat()
	dialog.Position = math.Vector{X: 50, Y: 300}

	return dialog
}
