package engine

type ITimeout interface {
	ILiving

	Update()
	Exec()
	GetId() string
}
