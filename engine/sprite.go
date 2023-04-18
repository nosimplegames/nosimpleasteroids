package engine

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type Sprite struct {
	math.Transformable
	Entity

	Texture render.Texture
}

func (sprite *Sprite) SetOriginCenter() {
	size := render.GetTextureSize(sprite.Texture)
	sprite.Origin = size.By(0.5)
}

func (sprite Sprite) Draw(target render.RenderTarget, transform math.Transform) {
	render.Sprite{
		Texture:   sprite.Texture,
		Target:    target,
		Transform: transform,
	}.Render()
}
