package animators

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type ISprite interface {
	SetRect(*math.Box)
	SetTexture(render.Texture)
}
