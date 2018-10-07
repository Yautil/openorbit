package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/ByteArena/box2d"
	"github.com/Yautil/engoBox2dSystem"
	"github.com/Yautil/openorbit/components"
)

type KeyboardMovementEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*engoBox2dSystem.Box2dComponent
	*components.SpeedInterpolationComponent
}

type KeyboardMovementSystem struct {
	entities []KeyboardMovementEntity
	world    *ecs.World
}

func (s *KeyboardMovementSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, box *engoBox2dSystem.Box2dComponent, speed *components.SpeedInterpolationComponent) {
	s.entities = append(s.entities, KeyboardMovementEntity{basic, space, box, speed})
}

// Remove removes the entity from the system
func (s *KeyboardMovementSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range s.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		s.entities = append(s.entities[:delete], s.entities[delete+1:]...)
	}
}

func (s *KeyboardMovementSystem) New(w *ecs.World) {
	engo.Input.RegisterAxis(
		"KeyboardMovementSystemVerticalAxis",
		engo.AxisKeyPair{engo.KeyArrowUp, engo.KeyArrowDown},
		engo.AxisKeyPair{engo.KeyW, engo.KeyS},
	)

	engo.Input.RegisterAxis(
		"KeyboardMovementSystemHorizontalAxis",
		engo.AxisKeyPair{engo.KeyArrowLeft, engo.KeyArrowRight},
		engo.AxisKeyPair{engo.KeyA, engo.KeyD},
	)
}

func (s *KeyboardMovementSystem) Update(dt float32) {
	for _, e := range s.entities {
		if e.SpaceComponent.Position.Y+e.SpaceComponent.Height > engo.GameHeight() {
			if engo.Input.Axis("KeyboardMovementSystemVerticalAxis").Value() > 0 {
				return
			}
		}

		if e.SpaceComponent.Position.X+e.SpaceComponent.Width > engo.GameWidth() {
			if engo.Input.Axis("KeyboardMovementSystemHorizontalAxis").Value() > 0 {
				return
			}
		}
		hori := engo.Input.Axis("KeyboardMovementSystemHorizontalAxis")
		//e.SpaceComponent.Position.X += e.Speed * hori.Value()

		vert := engo.Input.Axis("KeyboardMovementSystemVerticalAxis")
		//e.SpaceComponent.Position.Y += e.Speed * vert.Value()
		if hori.Value() != 0 || vert.Value() != 0 {
			e.SpeedInterpolationComponent.Moving = true
		} else {
			e.SpeedInterpolationComponent.Moving = false
			//continue
		}
		//e.Body.ApplyLinearImpulseToCenter(box2d.B2Vec2{X: float64(e.SpeedInterpolationComponent.CurSpeed * hori.Value()), Y: float64(e.SpeedInterpolationComponent.CurSpeed * vert.Value())}, true)
		e.Body.SetLinearVelocity(box2d.B2Vec2{X: float64(e.SpeedInterpolationComponent.CurSpeed * hori.Value()), Y: float64(e.SpeedInterpolationComponent.CurSpeed * vert.Value())})

	}
}
