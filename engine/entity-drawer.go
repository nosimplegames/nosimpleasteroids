package engine

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type EntityDrawer struct {
	Entity    IEntity
	Transform math.Transform
	Target    render.RenderTarget
}

func (drawer EntityDrawer) Draw() {
	canDraw := drawer.Entity.IsAlive() && drawer.Entity.IsVisible()
	if !canDraw {
		return
	}

	transform := drawer.Entity.GetTransform()
	transform.Concat(drawer.Transform)

	drawer.Entity.Draw(drawer.Target, transform)

	for _, child := range drawer.Entity.GetChildren() {
		EntityDrawer{
			Entity:    child,
			Transform: transform,
			Target:    drawer.Target,
		}.Draw()
	}
}
