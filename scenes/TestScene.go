package scenes

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/Yautil/engoBox2dSystem"
	"github.com/Yautil/openorbit/entities"
	"github.com/Yautil/openorbit/systems"
	"image/color"
)

type TestScene struct {
}

// Type uniquely defines your game type
func (*TestScene) Type() string { return "TestScene" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*TestScene) Preload() {
	engo.Files.Load("textures/asteroid.png")
	engo.Files.Load("textures/spaceship.png")
	engo.Files.Load("textures/star.png")
}

// Setup is called before the main loop starts. It allows you
// to add entities and systems to your Scene.
func (*TestScene) Setup(u engo.Updater) {
	// Config Scene World
	world, _ := u.(*ecs.World)
	common.SetBackground(color.Black)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.SpeedInterpolationSystem{})
	world.AddSystem(&systems.KeyboardMovementSystem{})
	world.AddSystem(&systems.LocalPlayerSystem{})
	world.AddSystem(&engoBox2dSystem.CollisionSystem{})
	world.AddSystem(&engoBox2dSystem.PhysicsSystem{})

	// Config Asteroid
	asteroid := entities.Star32x32Test{
		Name:        "Asteroid",
		Description: "uff",
		Texture:     "asteroid.png",
		Width:       64,
		Height:      64,
		World:       world,
		BasicEntity: ecs.NewBasic(),
	}
	asteroid.New(engo.Point{100, 200})

	star := entities.Star32x32Test{
		Name:        "Star",
		Description: "",
		Texture:     "star.png",
		Width:       32,
		Height:      32,
		World:       world,
	}
	star.New(engo.Point{300, 300})

	player := entities.LocalPlayer64x64Test{
		Name:        "Player 1",
		Description: "",
		Texture:     "spaceship.png",
		Width:       64,
		Height:      64,
		World:       world,
		BasicEntity: ecs.NewBasic(),
	}
	player.New(engo.Point{engo.GameWidth()/2 - float32(entities.LocalPlayer64x64Test{}.Width/2), engo.GameHeight()/2 - float32(entities.LocalPlayer64x64Test{}.Height/2)})
}
