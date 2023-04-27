package engine

type ILiving interface {
	IsAlive() bool
	IsRunning() bool
	Die()
}
