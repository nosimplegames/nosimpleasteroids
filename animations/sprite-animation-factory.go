package animations

import (
	"simple-games.com/asteroids/math"
	"simple-games.com/asteroids/render"
)

type SpriteAnimationFactory struct {
	Texture       render.Texture
	FrameDuration float64
	FrameSize     math.Vector
}

func (factory SpriteAnimationFactory) Create() SpriteAnimation {
	animation := SpriteAnimation{
		Animation: Animation{
			State: AnimationRunning,
		},

		Texture:       factory.Texture,
		FrameDuration: factory.FrameDuration,
		Frames:        factory.GetFrames(),
	}

	return animation
}

func (factory SpriteAnimationFactory) GetFrames() []math.Box {
	frames := []math.Box{}

	textureWidth, textureHeight := factory.Texture.Size()
	columns := textureWidth / int(factory.FrameSize.X)
	rows := textureHeight / int(factory.FrameSize.Y)
	halfFrameSize := factory.FrameSize.By(0.5)

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			frame := math.Box{
				Size: factory.FrameSize,
				Position: math.Vector{
					X: halfFrameSize.X + float64(column)*factory.FrameSize.X,
					Y: halfFrameSize.Y + float64(row)*factory.FrameSize.Y,
				},
			}
			frames = append(frames, frame)
		}
	}

	return frames
}
