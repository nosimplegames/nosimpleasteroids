package game

import (
	"golang.org/x/image/font"
	"simple-games.com/asteroids/animations"
	"simple-games.com/asteroids/animators"
	"simple-games.com/asteroids/assets"
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
	"simple-games.com/asteroids/res"
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
			SpaceshipTexture:      assets.LoadTexture(res.Spaceship),
			BulletTexture:         assets.LoadTexture(res.Bullet),
			BigAsteroidTexture:    assets.LoadTexture(res.BigAsteroid),
			NormalAsteroidTexture: assets.LoadTexture(res.NormalAsteroid),
			SmallAsteroidTexture:  assets.LoadTexture(res.SmallAsteroid),
			SmallExplosionTexture: assets.LoadTexture(res.SmallExplosion),
			BigExplosionTexture:   assets.LoadTexture(res.BigExplosion),
			EnergyBlastTexture:    assets.LoadTexture(res.EnergyBlast),
			ShieldTexture:         assets.LoadTexture(res.SimpleEnergyShield),

			BarTexture:        assets.LoadTexture(res.Bar),
			ProgressTexture:   assets.LoadTexture(res.Progress),
			NextDialogTexture: assets.LoadTexture(res.NextDialog),

			TextDialogSlice9: render.Slice9Factory{
				Texture: assets.LoadTexture(res.TextDialog),
				Top:     9,
				Right:   18,
				Bottom:  18,
				Left:    9,
			}.Create(),

			UIFontFace: assets.FontFaceFactory{
				Size: 24,
				DPI:  72,
				FontData: assets.FontData{
					Bytes: res.UIFont,
					Name:  "uifont",
				},
			}.Create(),

			TitleFontFace: assets.FontFaceFactory{
				Size: 180,
				DPI:  72,
				FontData: assets.FontData{
					Bytes: res.TitleFont,
					Name:  "titlefont",
				},
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
