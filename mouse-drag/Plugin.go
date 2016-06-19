package mouse_drag

import (
	"bellina"
	"xel"
	"fmt"
	"bellina/event"
)

var lastNodeID string
var startX, startY int32
var mouseOffsetX, mouseOffsetY int32

type Event struct {
	X, Y int32
	Target *bl.Node
	StartX, StartY int32
	MouseOffsetX, MouseOffsetY int32
}

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "mouse-drag"
}

func (c *Plugin) Tick() {
}

func (c *Plugin) Init() {
	event.RegisterLongTerm(bl.Mouse_Button_Event_Type, func(mouseButtonEvent event.Event) {

		e := mouseButtonEvent.(*bl.MouseButtonEvent)

		if e.Action == xel.Action_Up {
			lastNodeID = ""
		}
	})

	event.RegisterLongTerm(bl.Mouse_Move_Event_Type, func(mouseMoveEvent event.Event) {

		if lastNodeID == "" {
			return
		}

		e := mouseMoveEvent.(*bl.MouseMoveEvent)

		node := bl.GetNodeByID( lastNodeID )

		e.Target = node

		node.CallMouseMoveCallbacks(e)
	})
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On(cb func(interface{})) {
	c.On2(cb, nil, nil)
}

func (c *Plugin) On2(cb func(interface{}), startCb func(interface{}), endCb func(interface{})) {

	bl.OnMouseButton( func(e *bl.MouseButtonEvent) {
		if e.Action == xel.Action_Down {
			lastNodeID = e.Target.ID
			startX, startY = bl.Mouse_X, bl.Mouse_Y

			absX, absY := bl.GetNodeAbsolutePos(e.Target)
			mouseOffsetX = bl.Mouse_X - absX
			mouseOffsetY = bl.Mouse_Y - absY

			if startCb != nil {
				startCb(newEvent(bl.Mouse_X, bl.Mouse_Y, e.Target))
			}

		} else if e.Action == xel.Action_Up {
			lastNodeID = ""

			if endCb != nil {
				endCb(newEvent(bl.Mouse_X, bl.Mouse_Y, e.Target))
			}
		} else {
			fmt.Println("Button action not recognized in click.Plugin")
		}
	})

	bl.OnMouseMove( func(e *bl.MouseMoveEvent) {
		if lastNodeID == e.Target.ID {
			// we have a click!
			if cb != nil {
				cb(newEvent(e.X, e.Y, e.Target))
			}
		}
	})
}

func newEvent(mouseX, mouseY int32, target *bl.Node) Event {
	return Event{
		mouseX, mouseY,
		target,
		startX, startY,
		mouseOffsetX, mouseOffsetY,
	}
}

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}
