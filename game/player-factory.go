package game

import (
	"simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
)

type PlayerFactory struct {
	IsPlayerRespawning   bool
	IsControllerDisabled bool
}

func (factory PlayerFactory) Create() *Player {
	player := &Player{}
	player.Spaceship = *SpaceshipFactory{}.Create()
	player.IsControllerEnabled = !factory.IsControllerDisabled

	player.HealthBar = PlayerHealthBarFactory{}.Create()
	// player.AddChild(player.HealthBar)

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
