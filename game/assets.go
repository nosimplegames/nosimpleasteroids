package game

import (
	"golang.org/x/image/font"
	"simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/animators"
	"simple-games.com/asteroids/assets"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type Assets struct {
	SpaceshipTexture      render.Texture
	BulletTexture         render.Texture
	BigAsteroidTexture    render.Texture
	NormalAsteroidTexture render.Texture
	SmallAsteroidTexture  render.Texture
	SmallExplosionTexture render.Texture
	BigExplosionTexture   render.Texture
	EnergyBlastTexture    render.Texture
	ShieldTexture         render.Texture

	BarTexture        render.Texture
	ProgressTexture   render.Texture
	NextDialogTexture render.Texture
	TextDialogSlice9  render.Slice9

	UIFontFace    font.Face
	TitleFontFace font.Face

	FadeInAnimatorFactory          animators.AlphaAnimatorFactory
	TitleFadeInAnimatorFactory     animators.AlphaAnimatorFactory
	SmallExplosionAnimatorFactory  animators.SpriteAnimatorFactory
	BigExplosionAnimatorFactory    animators.SpriteAnimatorFactory
	SpaceshipShieldAnimatorFactory animators.SpriteAnimatorFactory
}

var globalAssets *Assets = nil

func GetAssets() *Assets {
	needToInitAssets := globalAssets == nil

	if needToInitAssets {
		globalAssets = &Assets{
			SpaceshipTexture:      assets.LoadTexture("./res/spaceship.png"),
			BulletTexture:         assets.LoadTexture("./res/bullet.png"),
			BigAsteroidTexture:    assets.LoadTexture("./res/big-asteroid.png"),
			NormalAsteroidTexture: assets.LoadTexture("./res/normal-asteroid.png"),
			SmallAsteroidTexture:  assets.LoadTexture("./res/small-asteroid.png"),
			SmallExplosionTexture: assets.LoadTexture("./res/small-explosion.png"),
			BigExplosionTexture:   assets.LoadTexture("./res/big-explosion.png"),
			EnergyBlastTexture:    assets.LoadTexture("./res/energy-blast.png"),
			ShieldTexture:         assets.LoadTexture("./res/simple-energy-shield.png"),

			BarTexture:        assets.LoadTexture("./res/bar.png"),
			ProgressTexture:   assets.LoadTexture("./res/progress.png"),
			NextDialogTexture: assets.LoadTexture("./res/next-dialog.png"),

			TextDialogSlice9: render.Slice9Factory{
				Texture: assets.LoadTexture("./res/text-dialog.png"),
				Top:     4,
				Right:   8,
				Bottom:  8,
				Left:    4,
			}.Create(),

			UIFontFace: assets.FontFaceFactory{
				Size:         16,
				DPI:          72,
				FontFileName: "./res/dogicapixel.ttf",
			}.Create(),

			TitleFontFace: assets.FontFaceFactory{
				Size:         60,
				DPI:          72,
				FontFileName: "./res/DayPosterShadowNF.ttf",
			}.Create(),
		}

		globalAssets.SmallExplosionAnimatorFactory = animators.SpriteAnimatorFactory{
			Animation: animations.SpriteAnimationFactory{
				Texture:       globalAssets.SmallExplosionTexture,
				FrameDuration: 1.0 / 48.0,
				FrameSize: math.Vector{
					X: 20,
					Y: 20,
				},
			}.Create(),
		}
		globalAssets.BigExplosionAnimatorFactory = animators.SpriteAnimatorFactory{
			Animation: animations.SpriteAnimationFactory{
				Texture:       globalAssets.BigExplosionTexture,
				FrameDuration: 1.0 / 24.0,
				FrameSize: math.Vector{
					X: 32,
					Y: 32,
				},
			}.Create(),
		}
		globalAssets.SpaceshipShieldAnimatorFactory = animators.SpriteAnimatorFactory{
			Animation: animations.SpriteAnimationFactory{
				Texture:       globalAssets.ShieldTexture,
				FrameDuration: 1.0 / 14.0,
				FrameSize: math.Vector{
					X: 32,
					Y: 32,
				},
				FrameCount: 7,
			}.Create(),
		}
		globalAssets.FadeInAnimatorFactory = animators.AlphaAnimatorFactory{
			Animation: animations.NumberAnimationFactory{
				InitialValue: 0,
				TargetValue:  1,
				Duration:     0.2,
			}.Create(),
		}
		globalAssets.TitleFadeInAnimatorFactory = animators.AlphaAnimatorFactory{
			Animation: animations.NumberAnimationFactory{
				InitialValue: 0,
				TargetValue:  1,
				Duration:     3,
			}.Create(),
		}
	}

	return globalAssets
}
