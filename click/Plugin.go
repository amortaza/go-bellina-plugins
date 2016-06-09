package click

import (
	"bellina"
	"xel"
	"fmt"
)

var lastNode *bl.Node

type Event struct {
	X, Y int32
	Target *bl.Node
}

type ClickPlugin struct {
}

func (c *ClickPlugin) Name() string {
	return "click"
}

func (c *ClickPlugin) Init() {
}

func (c *ClickPlugin) Uninit() {
}

func (c *ClickPlugin) On(cb func(interface{})) {

	bl.OnMouseButton( func(e *bl.MouseButtonEvent) {
		if e.Action == xel.Down {
			lastNode = e.Target
		} else if e.Action == xel.Up {
			if lastNode.ID == e.Target.ID {
				// we have a click!
				cb(Event{bl.Mouse_X, bl.Mouse_X, e.Target})
			}
		} else {
			fmt.Println("Button action not recognized in click.Plugin")
		}
	})
}

func NewPlugin() *ClickPlugin {
	c := &ClickPlugin{}

	return c
}
