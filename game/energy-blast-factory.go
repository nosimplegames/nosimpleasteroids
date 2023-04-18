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

	return blast
}

func (factory EnergyBlastFactory) GetTexture() render.Texture {
	targetSize := math.Vector{
		X: 1000,
		Y: 8,
	}
	texture := render.TextureFactory{
		TargetSize: targetSize,
	}.Create()

	render.TextureRepeater{
		Target:  texture,
		Texture: GetAssets().EnergyBlastTexture,
		Area: math.Box{
			Size: targetSize,
		},
	}.Repeat()

	return texture
}
