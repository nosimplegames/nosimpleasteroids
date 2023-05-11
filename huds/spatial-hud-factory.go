package huds

import (
	"simple-games.com/asteroids/assets"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/game"
	"simple-games.com/asteroids/ui"
)

type SpatialHudFactory struct {
	Player *game.Player
}

func (factory SpatialHudFactory) Create() engine.IEntity {
	hud := factory.GetHud()

	healthBar := game.PlayerHealthBarFactory{}.Create()
	hud.AddChild(healthBar)
	factory.Player.ConnectHealthBar(healthBar)

	return hud
}

func (factory SpatialHudFactory) GetHud() *ui.Container {
	hud := &ui.Container{}
	hud.Size = assets.GameSize
	hud.Padding = ui.SamePadding(assets.UIPadding)
	hud.SetPosition(assets.GameSize.By(0.5))
	hud.Layout = &ui.FlexLayout{
		LayoutDirection: ui.FlexRow,
	}

	return hud
}
