package engine

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type Sprite struct {
	Entity

	Rect    *math.Box
	Texture render.Texture
}

func (sprite *Sprite) SetRect(rect *math.Box) {
	sprite.Rect = rect
}

func (sprite *Sprite) SetTexture(texture render.Texture) {
	sprite.Texture = texture
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
		Rect:      sprite.Rect,
	}.Render()
}
