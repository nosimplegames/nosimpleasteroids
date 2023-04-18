package engine

type ITimeout interface {
	Update()
	Exec()
	IsAlive() bool
	Die()
	GetId() string
}
