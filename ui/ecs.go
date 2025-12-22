package ui

import ecs "github.com/igadmg/goecs"

type Context any

type LayoutComponentI interface {
	PrepareLayout() bool
	Layout(_lay *Context)
}

type LayoutComponent struct {
	ecs.MetaTag `ecs:"component"`

	Context
	Layout LayoutComponentI `gog:"new: 'layout'"`
}
