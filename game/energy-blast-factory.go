package game

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type EnergyBlastFactory struct {
}

func (factory EnergyBlastFactory) Create() *EnergyBlast {
	blast := &EnergyBlast{}
	blast.Texture = factory.GetTexture()
	blast.Inert.Origin = math.Vector{X: 16, Y: 400}
	blast.Inert.Position = math.Vector{X: 400, Y: 16}
	blast.Inert.Rotation = math.DegreesToRads(90)
	blast.Speed = 10

	return blast
}

func (factory EnergyBlastFactory) GetTexture() render.Texture {
	targetSize := math.Vector{
		X: 32,
		Y: 800,
	}
	texture := render.TextureFactory{
		TargetSize: targetSize,
	}.Create()

	render.TextureRepeater{
		Target:  texture,
		Texture: GetAssets().EnergyBlastTexture,
		Area: math.Box{
			Position: targetSize.By(0.5),
			Size:     targetSize,
		},
	}.Repeat()

	return texture
}
