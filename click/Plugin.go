package click

import (
	"bellina"
	"xel"
	"fmt"
)

var lastNodeID string

type Event struct {
	X, Y int32
	Target *bl.Node
}

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "click"
}

func (c *Plugin) Init() {
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not supoorted in click.Plugin")
}

func (c *Plugin) On(cb func(interface{})) {

	bl.OnMouseButton( func(e *bl.MouseButtonEvent) {
		if e.Action == xel.Down {
			lastNodeID = e.Target.ID
		} else if e.Action == xel.Up {
			if lastNodeID == e.Target.ID {
				// we have a click!
				cb(Event{bl.Mouse_X, bl.Mouse_X, e.Target})
			}
		} else {
			fmt.Println("Button action not recognized in click.Plugin")
		}
	})
}


func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}
