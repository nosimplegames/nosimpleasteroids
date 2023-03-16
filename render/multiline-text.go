package render

import (
	"golang.org/x/image/font"
	"simple-games.com/asteroids/math"
)

type MultilineText struct {
	Lines    []string
	FontFace font.Face
	Target   RenderTarget
	Position math.Vector
	Size     math.Vector
}

func (text MultilineText) Render() {
	renderingPosition := text.GetRenderingPosition()
	lineHeight := text.Size.Y / float64(len(text.Lines))

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

func (text MultilineText) GetRenderingPosition() math.Vector {
	return math.Vector{
		X: text.Position.X,
		Y: text.Position.Y - text.Size.Y*0.5,
	}
}
