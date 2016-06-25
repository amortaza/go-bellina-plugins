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

func (c *Plugin) Tick() {
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On(cb func(interface{})) {
	c.On2(cb, nil, nil)
}


func (c *Plugin) On2(cb func(interface{}), onDown func(interface{}), onUpAndMiss func(interface{})) {

	bl.OnMouseButton( func(e *bl.MouseButtonEvent) {

		if e.Action == xel.Button_Action_Down {
			lastNodeID = e.Target.ID

			if onDown != nil {
				onDown(Event{bl.Mouse_X, bl.Mouse_X, e.Target})
			}
		} else if e.Action == xel.Button_Action_Up {

			if lastNodeID == e.Target.ID {
				// we have a click!
				cb(Event{bl.Mouse_X, bl.Mouse_X, e.Target})

			} else {
				if onUpAndMiss != nil {
					onUpAndMiss(nil)
				}
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
