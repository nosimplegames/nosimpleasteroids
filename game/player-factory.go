package game

import (
	"simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
	"simple-games.com/asteroids/math"
)

type PlayerFactory struct {
	IsPlayerRespawning bool
}

func (factory PlayerFactory) Create() *Player {
	player := &Player{}
	player.Propulsor = Engine{
		MinSpeed:      0.3,
		Acceleration:  0.05,
		MaxSpeed:      4,
		Deaceleration: 0.025,
	}
	player.RightRotator = Engine{
		MinSpeed:      0.0,
		Acceleration:  0.003,
		MaxSpeed:      0.05,
		Deaceleration: 0.002,
	}
	player.LeftRotator = Engine{
		MinSpeed:      0.0,
		Acceleration:  0.003,
		MaxSpeed:      0.05,
		Deaceleration: 0.002,
	}
	player.Turret = &Turret{
		Weapon: Weapon{
			TimeBetweenShoots: 0.1,
		},
	}
	player.LifePoints = 3

	player.Position = math.Vector{
		X: 400,
		Y: 400,
	}
	player.Origin = math.Vector{
		X: 8,
		Y: 8,
	}

	player.HealthBar = PlayerHealthBarFactory{}.Create()
	player.AddChild(player.HealthBar)

	player.AddEventListener(events.EventListener{
		EventType: LifePointsChanged,
		Callback: func(_ events.Event) {
			player.HealthBar.Value = float64(player.LifePoints)
		},
	})

	// Theres a problem here, once animation is done, it'll be removed from
	// engine animations but player will still has a pointer to it,
	// even though it will be returning target value, player pointer may be null
	if factory.IsPlayerRespawning {
		animation := factory.GetRespawiningAnimation()
		player.RespawningAnimation = animation
		engine.GetAnimations().AddAnimation(animation)
	}

	player.OnDestroy = func() {
		engine.GetTimer().AddTimeout(&engine.Timeout{
			Callback: func() {
				player := PlayerFactory{
					IsPlayerRespawning: true,
				}.Create()
				engine.GetEntities().AddEntity(player)
			},
			Time: 1.5,
		})

		ExplosionFactory{
			Position:      player.Position,
			ExplosionSize: BigExplosion,
		}.Create()
	}

	engine.GetWorld().AddCollisinable(player)
	return player
}

func (factory PlayerFactory) GetRespawiningAnimation() *animations.AnimationList {
	animation := GetAssets().RespawningAnimation.Copy()

	return animation.(*animations.AnimationList)
}
