package render

import (
	"golang.org/x/image/font"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/utils/utilstext"
)

type MultilineText struct {
	Lines         []string
	FontFace      font.Face
	LineHeight    float64
	Target        RenderTarget
	Position      math.Vector
	TextAlignment TextAlignment
}

func (text MultilineText) Render() {
	linesRenderingPositions := text.GetLinesRenderingPositions()

	for lineIndex, line := range text.Lines {
		position := linesRenderingPositions[lineIndex]

		Text{
			Text:     line,
			Target:   text.Target,
			FontFace: text.FontFace,
			Position: position,
		}.Render()
	}
}

func (text MultilineText) GetLinesRenderingPositions() []math.Vector {
	lineFontHeight := utilstext.GetFontFaceHeight(text.FontFace)
	lineHeight := lineFontHeight * text.LineHeight
	renderingPosition := text.GetRenderingPosition(lineHeight)
	positions := []math.Vector{}
	lineTopOffset := (lineHeight - lineFontHeight) / 2.0

	for lineIndex := range text.Lines {
		position := math.Vector{
			X: renderingPosition.X,
			Y: renderingPosition.Y + lineTopOffset + lineHeight*float64(lineIndex),
		}
		positions = append(positions, position)
	}

	return positions
}

func (text MultilineText) GetRenderingPosition(lineHeight float64) math.Vector {
	return math.Vector{
		X: text.Position.X,
		Y: text.Position.Y - lineHeight,
	}
}
