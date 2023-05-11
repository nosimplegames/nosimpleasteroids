package game

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type EnergyBlastFactory struct {
}

func (factory EnergyBlastFactory) Create() *EnergyBlast {
	blast := &EnergyBlast{}
	blast.Texture = factory.GetTexture()
	blast.SetOrigin(math.Vector{X: 16, Y: 400})
	blast.SetPosition(math.Vector{X: 400, Y: 16})
	blast.SetRotation(math.DegreesToRads(90))
	blast.Speed = 10
	// Size depends on rotation as we're using AABB
	blast.Size = math.Vector{X: 800, Y: 32}

	engine.GetWorld().AddCollisinable(blast)

	return blast
}

func (factory EnergyBlastFactory) GetTexture() render.Texture {
	targetSize := factory.GetTargetSize()
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

func (factory EnergyBlastFactory) GetTargetSize() math.Vector {
	return math.Vector{
		X: 32,
		Y: 800,
	}
}
