package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/Yautil/engoBox2dSystem"
	"log"
)

type LocalPlayerEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	name string
}

type LocalPlayerSystem struct {
	entities []LocalPlayerEntity
	world    *ecs.World
	score    int
}

func (s *LocalPlayerSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent, name string) {
	s.entities = append(s.entities, LocalPlayerEntity{basic, space, name})
}

// Remove removes the entity from the system
func (s *LocalPlayerSystem) Remove(basic ecs.BasicEntity) {
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

func (s *LocalPlayerSystem) New(w *ecs.World) {

	engo.Mailbox.Listen("CollisionStartMessage", func(message engo.Message) {
		for _, e := range s.entities {
			c, isCollision := message.(engoBox2dSystem.CollisionStartMessage)
			if isCollision {
				if c.Contact.IsTouching() {
					a := c.Contact.GetFixtureA().GetBody().M_userData
					b := c.Contact.GetFixtureB().GetBody().GetUserData()
					log.Println(e.name)
					if a == e.name || b == e.name {
						if a == "Asteroid" || b == "Asteroid" {
							log.Println("DIED")
						}
						if a == "Star" || b == "Star" {
							log.Println("SCORED")
						}
					}
				}
			}
		}
	})
}

func (s *LocalPlayerSystem) Update(dt float32) {
}
