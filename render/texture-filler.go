package render

import "simple-games.com/asteroids/math"

type TextureFiller struct {
	Target  RenderTarget
	Texture Texture

	Columns          int
	Rows             int
	StartingPosition math.Vector
}

func (filler TextureFiller) Render() {
	textureSize := GetTextureSize(filler.Texture)

	for row := 0; row < filler.Rows; row++ {
		textureY := filler.StartingPosition.Y + float64(row)*textureSize.Y

		for column := 0; column < filler.Columns; column++ {
			textureX := filler.StartingPosition.X + float64(column)*textureSize.X

			position := math.Vector{
				X: textureX,
				Y: textureY,
			}

			transform := math.Transformable{Position: position}.GetTransform()

			Sprite{
				Target:    filler.Target,
				Texture:   filler.Texture,
				Transform: transform,
			}.Render()
		}
	}
}
