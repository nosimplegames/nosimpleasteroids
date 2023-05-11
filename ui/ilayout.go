package ui

import "simple-games.com/asteroids/math"

type ILayout interface {
	GetPositions(elements []LayoutElement) []math.Vector
	SetSize(math.Vector)
}
