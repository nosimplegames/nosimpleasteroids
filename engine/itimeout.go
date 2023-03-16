package engine

type ITimeout interface {
	Update()
	Exec()
	IsAlive() bool
}
