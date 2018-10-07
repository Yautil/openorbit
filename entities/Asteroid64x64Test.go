package entities

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/ByteArena/box2d"
	"github.com/Yautil/engoBox2dSystem"
	"log"
)

type Asteroid64x64Test struct {
	// Unified Entity Information
	Name        string
	Description string
	Texture     string
	Width       int
	Height      int
	World       *ecs.World
	// Components
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	engoBox2dSystem.Box2dComponent
	// Specific Entries
}

func (o *Asteroid64x64Test) New(position engo.Point) {
	e := Asteroid64x64Test{BasicEntity: ecs.NewBasic()}
	e.SpaceComponent = common.SpaceComponent{
		Position: position,
		Width:    float32(o.Width),
		Height:   float32(o.Height),
	}
	t, err := common.LoadedSprite("textures/" + o.Texture)
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
		e.RenderComponent = common.RenderComponent{
			Scale: engo.Point{1, 1},
		}
	} else {
		e.RenderComponent = common.RenderComponent{
			Scale:    engo.Point{1, 1},
			Drawable: t,
		}
	}
	//box2d component setup
	body := box2d.NewB2BodyDef()
	body.Type = box2d.B2BodyType.B2_dynamicBody
	body.Position = engoBox2dSystem.Conv.ToBox2d2Vec(e.Center())
	body.Angle = engoBox2dSystem.Conv.DegToRad(e.Rotation)
	body.UserData = o.Name
	e.Box2dComponent.Body = engoBox2dSystem.World.CreateBody(body)

	var shape box2d.B2PolygonShape
	shape.SetAsBox(engoBox2dSystem.Conv.PxToMeters(e.SpaceComponent.Width/2),
		engoBox2dSystem.Conv.PxToMeters(e.SpaceComponent.Height/2))
	dudeFixtureDef := box2d.B2FixtureDef{
		Shape:    &shape,
		Density:  1.0,
		Friction: 0.1,
	}
	e.Box2dComponent.Body.CreateFixtureFromDef(&dudeFixtureDef)
	for _, system := range o.World.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.AddByInterface(&e)
		case *engoBox2dSystem.PhysicsSystem:
			sys.Add(&e.BasicEntity, &e.SpaceComponent, &e.Box2dComponent)
		case *engoBox2dSystem.CollisionSystem:
			sys.AddByInterface(&e)
		}
	}
	log.Println("Generated new Entity: " + o.Name)
}
