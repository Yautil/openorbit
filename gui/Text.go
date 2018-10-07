package gui

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type Text struct {
	Font     *common.Font
	cache    string
	Position engo.Point
	World    *ecs.World
	Text     string

	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (t *Text) Draw() {
	t.BasicEntity = ecs.NewBasic()

	width, height, _ := t.Font.TextDimensions(t.Text)

	t.SpaceComponent = common.SpaceComponent{
		Width:  float32(width),
		Height: float32(height),
	}
	t.SpaceComponent.Position = t.Position
	t.RenderComponent.Drawable = common.Text{
		Font: t.Font,
		Text: t.Text,
	}

	t.SetShader(common.TextHUDShader)

	for _, system := range t.World.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
		}
	}
}

func (t *Text) SetText(s string) bool {
	if t.Font == nil {
		panic("Text.SetText called without setting Text.Font")
	}

	if t.cache == s {
		return false
	}

	if t.RenderComponent.Drawable == nil {
		t.RenderComponent.Drawable = common.Text{Font: t.Font}
	}

	fnt := t.RenderComponent.Drawable.(common.Text)
	fnt.Text = s

	return true
}
