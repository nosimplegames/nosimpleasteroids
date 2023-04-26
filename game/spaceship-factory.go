package game

import "simple-games.com/asteroids/math"

type SpaceshipFactory struct {
}

func (factory SpaceshipFactory) Create() *Spaceship {
	spaceship := &Spaceship{}

	spaceship.Texture = GetAssets().SpaceshipTexture

	spaceship.Propulsor = Engine{
		MinSpeed:      0.3,
		Acceleration:  0.05,
		MaxSpeed:      4,
		Deaceleration: 0.025,
	}
	spaceship.RightRotator = Engine{
		MinSpeed:      0.0,
		Acceleration:  0.003,
		MaxSpeed:      0.05,
		Deaceleration: 0.002,
	}
	spaceship.LeftRotator = Engine{
		MinSpeed:      0.0,
		Acceleration:  0.003,
		MaxSpeed:      0.05,
		Deaceleration: 0.002,
	}
	spaceship.Turret = &Turret{
		Weapon: Weapon{
			TimeBetweenShoots: 0.1,
		},
	}
	spaceship.LifePoints = 3

	spaceship.Position = math.Vector{
		X: 400,
		Y: 400,
	}
	spaceship.Origin = math.Vector{
		X: 8,
		Y: 8,
	}
	spaceship.Rotation = math.DegreesToRads(-90)

	return spaceship
}
