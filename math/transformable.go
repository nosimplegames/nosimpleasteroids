package math

type Transformable struct {
	Position Vector
	Origin   Vector
	Rotation float64
}

func (transformable *Transformable) Move(movementVector Vector) {
	transformable.Position.X += movementVector.X
	transformable.Position.Y += movementVector.Y
}

func (transformable *Transformable) Rotate(tetha float64) {
	transformable.Rotation += tetha
}

func (transformable Transformable) GetTransform() Transform {
	transform := Transform{}

	transform.Translate(
		-transformable.Origin.X,
		-transformable.Origin.Y,
	)
	transform.Rotate(transformable.Rotation)
	transform.Translate(
		transformable.Position.X,
		transformable.Position.Y,
	)

	return transform
}

func (transformable Transformable) GetPosition() Vector {
	return transformable.Position
}
