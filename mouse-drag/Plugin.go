package mouse_drag

import (
	"bellina"
	"xel"
	"fmt"
	"bellina/event"
)

var lastNodeID string

type Event struct {
	X, Y int32
	Target *bl.Node
}

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "mouse-drag"
}

func (c *Plugin) Init() {
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On(cb func(interface{})) {

	event.RegisterShortTerm(bl.Mouse_Button_Event_Type, func(mouseButtonEvent event.Event) {

		e := mouseButtonEvent.(*bl.MouseButtonEvent)

		if e.Action == xel.Up {
			lastNodeID = ""
		}
	})

	bl.OnMouseButton( func(e *bl.MouseButtonEvent) {
		if e.Action == xel.Down {
			lastNodeID = e.Target.ID

		} else if e.Action == xel.Up {
			lastNodeID = ""

		} else {
			fmt.Println("Button action not recognized in click.Plugin")
		}
	})

	event.RegisterShortTerm(bl.Mouse_Move_Event_Type, func(mouseMoveEvent event.Event) {

		if lastNodeID == "" {
			return
		}

		e := mouseMoveEvent.(*bl.MouseMoveEvent)

		node := bl.GetNodeByID( lastNodeID )

		e.Target = node

		node.CallMouseMoveCallbacks(e)
	})

	bl.OnMouseMove( func(e *bl.MouseMoveEvent) {
		if lastNodeID == e.Target.ID {
			// we have a click!
			cb(Event{bl.Mouse_X, bl.Mouse_X, e.Target})
		}
	})
}


func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}
