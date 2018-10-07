package systems

import (
	"engo.io/ecs"
	"github.com/Yautil/openorbit/components"
)

type SpeedInterpolationEntity struct {
	*ecs.BasicEntity
	*components.SpeedInterpolationComponent
}

type SpeedInterpolationSystem struct {
	entities []SpeedInterpolationEntity
	world    *ecs.World
}

func (s *SpeedInterpolationSystem) Add(basic *ecs.BasicEntity, speed *components.SpeedInterpolationComponent) {
	s.entities = append(s.entities, SpeedInterpolationEntity{basic, speed})
}

// Remove removes the entity from the system
func (s *SpeedInterpolationSystem) Remove(basic ecs.BasicEntity) {
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

func (s *SpeedInterpolationSystem) New(w *ecs.World) {

}

func (s *SpeedInterpolationSystem) Update(dt float32) {
	for _, e := range s.entities {
		if e.SpeedInterpolationComponent.Moving {
			if e.SpeedInterpolationComponent.CurSpeed < e.SpeedInterpolationComponent.MaxSpeed {
				e.SpeedInterpolationComponent.CurSpeed = e.SpeedInterpolationComponent.CurSpeed + e.SpeedInterpolationComponent.InterpolationDelta
			}
			if e.SpeedInterpolationComponent.CurSpeed > e.SpeedInterpolationComponent.MaxSpeed {
				e.SpeedInterpolationComponent.CurSpeed = e.SpeedInterpolationComponent.MaxSpeed
			}
		} else if e.SpeedInterpolationComponent.CurSpeed > 0 {
			e.SpeedInterpolationComponent.CurSpeed = e.SpeedInterpolationComponent.CurSpeed - e.SpeedInterpolationComponent.InterpolationDelta*e.SpeedInterpolationComponent.SlowdownMultiplier
		}
		if e.SpeedInterpolationComponent.CurSpeed < 0 {
			e.SpeedInterpolationComponent.CurSpeed = 0
		}
	}
}
