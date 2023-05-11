package asteroidsscene

import (
	"simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/engine"
	"simple-games.com/asteroids/game"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/scenes"
)

type GameTitle struct {
	scenes.Act
}

func (act *GameTitle) Openning(scene *scenes.Scene) {
	engine.EntityRemover{
		Parent: scene,
		Id:     "player",
	}.Remove()

	title := act.GetTitle()
	animation := act.GetTitleAnimation(title)

	engine.GetAnimations().AddAnimation(animation)
	act.AddChild(title)

	animation.OnStop.AddCallback(func() {
		act.Die()
	})
}

func (act GameTitle) GetTitle() *engine.MultilineText {
	// This may be moved to a GameTitle entity
	text := &engine.MultilineText{
		Lines: []string{
			"No",
			"Asteroids",
		},
		LineHeight: 1.5,
		FontFace:   game.GetAssets().TitleFontFace,
	}
	text.SetPosition(math.Vector{
		X: 400,
		Y: 400,
	})

	return text
}

func (act GameTitle) GetTitleAnimation(text *engine.MultilineText) *animations.AnimationList {
	fadeInAnimator := game.GetAssets().TitleFadeInAnimatorFactory.Create(text)

	animation := animations.AnimationLoopFactory{
		Animation:             fadeInAnimator,
		LoopCount:             1,
		IncludeBackwards:      true,
		TimeBetweenAnimations: 1,
	}.Create()

	return &animation
}
