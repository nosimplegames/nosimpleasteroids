package game

import "simple-games.com/asteroids/math"

type IWeapon interface {
	Shoot(position math.Vector, rotation float64)
	PrepareForNextShoot()
}
