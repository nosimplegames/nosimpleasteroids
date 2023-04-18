package render

import (
	"github.com/hajimehoshi/ebiten"
	"simple-games.com/asteroids/math"
)

type TextureFactory struct {
	TargetSize math.Vector
}

func (factory TextureFactory) Create() Texture {
	texture, _ := ebiten.NewImage(
		int(factory.TargetSize.X),
		int(factory.TargetSize.Y),
		ebiten.FilterDefault,
	)

	return texture
}
