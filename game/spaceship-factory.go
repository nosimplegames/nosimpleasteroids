package game

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type SpaceshipFactory struct {
}

func (factory SpaceshipFactory) Create() *Spaceship {
	spaceship := &Spaceship{}
	spaceship.Type = SpaceshipEntityType

	spaceship.Texture = GetAssets().SpaceshipTexture
	spaceship.Size = math.Vector{
		X: 39,
		Y: 39,
	}

	spaceship.Propulsor = Engine{
		MinSpeed:      3,
		Acceleration:  0.05,
		MaxSpeed:      6,
		Deaceleration: 0.01,
	}
	spaceship.RightRotator = Engine{
		MinSpeed:      0.0,
		Acceleration:  0.02,
		MaxSpeed:      0.08,
		Deaceleration: 0.005,
	}
	spaceship.LeftRotator = Engine{
		MinSpeed:      0.0,
		Acceleration:  0.02,
		MaxSpeed:      0.08,
		Deaceleration: 0.005,
	}
	spaceship.Turret = &Turret{
		Weapon: Weapon{
			TimeBetweenShoots: 0.1,
		},
	}

	spaceship.SetMaxLifePoints(100)
	spaceship.SetLifePoints(100)

	spaceship.SetOrigin(render.GetTextureSize(spaceship.Texture).By(0.5))
	spaceship.SetRotation(math.DegreesToRads(-90))

	return spaceship
}
