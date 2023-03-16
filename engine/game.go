package engine

import "github.com/hajimehoshi/ebiten"

type Game struct {
}

func (game *Game) Update(*ebiten.Image) error {
	GetUpdatables().Update()
	GetEntities().Update()
	GetWorld().Update()
	GetAnimations().Update()

	return nil
}

func (game Game) Draw(screen *ebiten.Image) {
	GetEntities().Draw(screen)
}

func (game Game) Layout(_, _ int) (int, int) {
	return 800, 800
}

func (game *Game) Run() {
	ebiten.SetWindowSize(800, 800)
	ebiten.RunGame(game)
}
