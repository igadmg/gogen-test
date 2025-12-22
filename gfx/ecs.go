package gfx

import (
	"github.com/igadmg/gamemath/rect2"
	ecs "github.com/igadmg/goecs"
)

type DrawCallFn func(source rect2.Float32)

type DrawComponent struct {
	ecs.MetaTag `ecs:"component"`

	DrawCall DrawCallFn
}

type DrawEntity struct {
	ecs.MetaTag `ecs:"archetype: { transient }"`
	Id          ecs.Id

	Draw *DrawComponent `gog:"new: drawCall"`
}

type DoubleDrawEntity struct {
	ecs.MetaTag `ecs:"archetype: { transient }"`
	DrawEntity  `gog:"new"`

	DrawAgain *DrawComponent `gog:"new: drawCall"`
}

type BoundComponent struct {
	ecs.MetaTag `ecs:"component"`

	Bound rect2.Float32
}

type BoundDrawEntity struct {
	ecs.MetaTag `ecs:"archetype: { transient }"`
	DrawEntity  `gog:"new: drawCall"`

	Bound BoundComponent `gog:"new: bound"`
}

// Here we removed base class new parameters and expect them to be imported automaticaly
type BoundDrawEntitySimple struct {
	ecs.MetaTag `ecs:"archetype: { transient }"`
	DrawEntity  `gog:"new"`

	Bound BoundComponent `gog:"new: bound"`
}
