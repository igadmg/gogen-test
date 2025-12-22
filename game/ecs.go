package game

import (
	ecs "github.com/igadmg/goecs"
	"github.com/igadmg/gogen-test/gfx"
	"github.com/igadmg/gogen-test/ui"
)

type WorldModelComponent struct {
	ecs.MetaTag `ecs:"component"`

	Date string `ecs:"a, dto"`
}

type WorldEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	Id          ecs.Id

	Model *WorldModelComponent
}

type ScreenLayoutComponent struct {
	ecs.MetaTag        `ecs:"component: { transient }"`
	ui.LayoutComponent `gog:"new"`
}

type ScreenViewComponent struct {
	ecs.MetaTag `ecs:"component"`

	background ecs.Ref[gfx.DrawEntity] `ecs:"a"` // background is transient ref here, because DrawEntity s transient, should be created automaticaly on create transient step
}

type ScreenModelComponent struct {
	ecs.MetaTag `ecs:"component"`

	worldMap ecs.Ref[WorldEntity] `ecs:"a, reference"` // reference components should not be created by default, but also not recreated as transient refs
}

type ScreenEntity struct {
	ecs.MetaTag `ecs:"archetype"`
	Id          ecs.Id

	layout *ScreenLayoutComponent `gog:"new: '@'"`
	View   *ScreenViewComponent   `gog:"new: { background: '@.DrawBackground' }"`
	Model  *ScreenModelComponent
}

func (s ScreenEntity) DrawBackground() {

}
