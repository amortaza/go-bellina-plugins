package mouse_drag

import (
	"github.com/amortaza/go-xel"
	"fmt"
	"github.com/amortaza/go-bellina/event"
	"github.com/amortaza/go-bellina"
)

var plugin *Plugin

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "mouse-drag"
}

func (c *Plugin) GetState() interface{} {
	return nil
}

func (c *Plugin) Tick() {
}

func (c *Plugin) Reset() {
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Init() {
	event.RegisterLongTerm(bl.EventType_Mouse_Button, func(mouseButtonEvent event.Event) {

		e := mouseButtonEvent.(*bl.MouseButtonEvent)

		if e.ButtonAction == xel.Button_Action_Up {
			gLastNodeId = ""
		}
	})

	event.RegisterLongTerm(bl.EventType_Mouse_Move, func(mouseMoveEvent event.Event) {

		if gLastNodeId == "" {
			return
		}

		e := mouseMoveEvent.(*bl.MouseMoveEvent)

		node := bl.GetNodeById(gLastNodeId)

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
		if e.ButtonAction == xel.Button_Action_Down {
			gLastNodeId = e.Target.Id
			startX, startY = bl.Mouse_X, bl.Mouse_Y

			absX, absY := bl.GetNodeAbsolutePos(e.Target)
			mouseOffsetX = bl.Mouse_X - absX
			mouseOffsetY = bl.Mouse_Y - absY

			if startCb != nil {
				startCb(newEvent(bl.Mouse_X, bl.Mouse_Y, e.Target))
			}

		} else if e.ButtonAction == xel.Button_Action_Up {
			gLastNodeId = ""

			if endCb != nil {
				endCb(newEvent(bl.Mouse_X, bl.Mouse_Y, e.Target))
			}
		} else {
			fmt.Println("Button action not recognized in click.Plugin")
		}
	})

	bl.OnMouseMove( func(e *bl.MouseMoveEvent) {
		if gLastNodeId == e.Target.Id {
			fmt.Println("Getting drag on " + gLastNodeId)
			// we have a drag!
			if cb != nil {
				cb(newEvent(e.X, e.Y, e.Target))
			}
		}
	})
}


