package asteroidsscene

import (
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/events"
	"simple-games.com/asteroids/game"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/scenes"
)

type Asteroids struct {
	scenes.Act

	Player *game.Player

	hasShownTeleportDialog bool
}

func (act *Asteroids) Openning(_ *scenes.Scene) {
	act.Player.IsControllerEnabled = true
	act.CreateTeleports()

	generator := game.GetAsteroidsGenerator()
	generator.OnDie.AddCallback(act.Die)

	act.Player.OnDie.AddCallback(generator.DieAfterAsteroids)
}

func (act *Asteroids) CreateTeleports() {
	targetPosition := math.Vector{
		X: 400,
		Y: 400,
	}

	levelSize := math.Vector{
		X: 800,
		Y: 800,
	}

	horizontalTeleportSize := math.Vector{
		X: 800,
		Y: 20,
	}
	verticalTeleportSize := math.Vector{
		X: 20,
		Y: 800,
	}
	teleportOffset := 10.0
	teleportCount := 4

	teleportSizes := []math.Vector{
		horizontalTeleportSize,
		verticalTeleportSize,
		horizontalTeleportSize,
		verticalTeleportSize,
	}

	teleportPositions := []math.Vector{
		{
			X: levelSize.X * 0.5,
			Y: -teleportOffset,
		},
		{
			X: levelSize.X + teleportOffset,
			Y: levelSize.Y * 0.5,
		},
		{
			X: levelSize.X * 0.5,
			Y: levelSize.Y + teleportOffset,
		},
		{
			X: -teleportOffset,
			Y: levelSize.Y * 0.5,
		},
	}

	teleports := []*game.Teleport{}

	for i := 0; i < teleportCount; i++ {
		teleport := &game.Teleport{
			CollisionMask:       game.SpaceshipTeleportCollisionMask,
			Size:                teleportSizes[i],
			Position:            teleportPositions[i],
			Target:              act.Player,
			TargetPosition:      targetPosition,
			TargetCollisionMask: act.Player.GetCollisionMask(),
		}

		teleports = append(teleports, teleport)
		engine.GetWorld().AddCollisinable(teleport)

		teleport.OnCollisionWithTarget.AddTCallback(act.ShowTeleportCollisionConversation)
	}

	act.OnDie(func() {
		for _, teleport := range teleports {
			teleport.Die()
		}
	})
}

func (act *Asteroids) ShowTeleportCollisionConversation(data events.SignalData) {
	teleport := data.(*game.Teleport)

	if act.hasShownTeleportDialog {
		teleport.Teleport()
		return
	}

	act.Player.Pause()
	act.Player.Hide()

	conversation := game.TextDialogFactory{
		Conversation: "For some reason, I just can't escape from here. May I fight until death!?",
		OnEnd: func() {
			teleport.Teleport()
			act.Player.Show()
			act.Player.Resume()
		},
	}.Create()

	act.AddChild(conversation)
	act.hasShownTeleportDialog = true
}
