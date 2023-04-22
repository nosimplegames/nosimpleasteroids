package textdialog

import (
	"math"

	"golang.org/x/image/font"
	"simple-games.com/asteroids/render"
	"simple-games.com/asteroids/utils/utilstext"
)

type DialogsFactory struct {
	TextLinesCount int
	FontFace       font.Face
	LineWidth      float64
	LineHeight     float64
	Text           string
}

func (factory DialogsFactory) Create() []Dialog {
	textLines := factory.GetTextLines()
	dialogs := []Dialog{}

	totalDialogs := int(math.Ceil(float64(len(textLines)) / float64(factory.TextLinesCount)))
	for i := 0; i < totalDialogs; i++ {
		leftIndex := i * factory.TextLinesCount
		rightIndex := leftIndex + factory.TextLinesCount
		isRightIndexOutOfRange := rightIndex > len(textLines)

		if isRightIndexOutOfRange {
			rightIndex = len(textLines)
		}

		dialog := Dialog{
			Lines: textLines[leftIndex:rightIndex],
		}
		dialogs = append(dialogs, dialog)
	}

	return dialogs
}

func (factory DialogsFactory) GetTextLines() []string {
	fontHeight := utilstext.GetFontFaceHeight(factory.FontFace)
	lineHeight := fontHeight * factory.LineHeight

	textLines, _ := render.TextLinesCalculator{
		Text:       factory.Text,
		LineHeight: lineHeight,
		LineWidth:  factory.LineWidth,
		FontFace:   factory.FontFace,
	}.Calculate()

	return textLines
}
