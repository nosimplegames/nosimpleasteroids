package ui

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/math"
)

type Container struct {
	engine.Entity

	Padding Padding
	Layout  ILayout
}

func (container *Container) Update() {
	hasLayout := container.Layout != nil

	if !hasLayout {
		return
	}

	container.RepositionElements()
}

func (container *Container) RepositionElements() {
	positions := container.Layout.GetPositions(container.GetElements())
	container.Layout.SetSize(container.GetLayoutSize())

	for childIndex, child := range container.Children {
		position := positions[childIndex]
		child.SetPosition(position)
	}
}

func (container *Container) GetElements() []LayoutElement {
	elements := []LayoutElement{}

	for _, child := range container.Children {
		elements = append(elements, LayoutElement{
			Size: child.GetSize(),
		})
	}

	return elements
}

func (container Container) GetLayoutSize() math.Vector {
	layoutSize := math.Vector{
		X: container.Size.X - container.Padding.Right - container.Padding.Left,
		Y: container.Size.Y - container.Padding.Top - container.Padding.Bottom,
	}

	return layoutSize
}
