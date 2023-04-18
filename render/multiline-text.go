package render

import (
	"golang.org/x/image/font"
	"simple-games.com/asteroids/math"
)

type MultilineText struct {
	Lines         []string
	FontFace      font.Face
	Target        RenderTarget
	Position      math.Vector
	TextAlignment TextAlignment
}

func (text MultilineText) Render() {
	lineHeight := float64(text.FontFace.Metrics().Height.Round())
	renderingPosition := text.GetRenderingPosition(lineHeight)

	for lineIndex, line := range text.Lines {
		position := math.Vector{
			X: renderingPosition.X,
			Y: renderingPosition.Y + lineHeight*float64(lineIndex),
		}

		Text{
			Text:     line,
			Target:   text.Target,
			FontFace: text.FontFace,
			Position: position,
		}.Render()
	}
}

func (text MultilineText) GetRenderingPosition(lineHeight float64) math.Vector {
	return math.Vector{
		X: text.Position.X,
		Y: text.Position.Y - lineHeight,
	}
}
