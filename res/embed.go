package res

import (
	_ "embed"
)

var (
	//go:embed spaceship.png
	Spaceship []byte

	//go:embed bullet.png
	Bullet []byte

	//go:embed big-asteroid.png
	BigAsteroid []byte

	//go:embed normal-asteroid.png
	NormalAsteroid []byte

	//go:embed small-asteroid.png
	SmallAsteroid []byte

	//go:embed small-explosion.png
	SmallExplosion []byte

	//go:embed big-explosion.png
	BigExplosion []byte

	//go:embed energy-blast.png
	EnergyBlast []byte

	//go:embed simple-energy-shield.png
	SimpleEnergyShield []byte

	//go:embed bar.png
	Bar []byte

	//go:embed progress.png
	Progress []byte

	//go:embed next-dialog.png
	NextDialog []byte

	//go:embed text-dialog.png
	TextDialog []byte

	//go:embed dogicapixel.ttf
	UIFont []byte

	//go:embed DayPosterShadowNF.ttf
	TitleFont []byte
)
