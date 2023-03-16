package animations

type IAnimation interface {
	Update()
	IsAlive() bool
	Copy() IAnimation
	Reverse() IAnimation
}
