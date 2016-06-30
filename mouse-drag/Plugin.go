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
			lastNodeID = ""
		}
	})

	event.RegisterLongTerm(bl.EventType_Mouse_Move, func(mouseMoveEvent event.Event) {

		if lastNodeID == "" {
			return
		}

		e := mouseMoveEvent.(*bl.MouseMoveEvent)

		node := bl.GetNodeById( lastNodeID )

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
			lastNodeID = e.Target.Id
			startX, startY = bl.Mouse_X, bl.Mouse_Y

			absX, absY := bl.GetNodeAbsolutePos(e.Target)
			mouseOffsetX = bl.Mouse_X - absX
			mouseOffsetY = bl.Mouse_Y - absY

			if startCb != nil {
				startCb(newEvent(bl.Mouse_X, bl.Mouse_Y, e.Target))
			}

		} else if e.ButtonAction == xel.Button_Action_Up {
			lastNodeID = ""

			if endCb != nil {
				endCb(newEvent(bl.Mouse_X, bl.Mouse_Y, e.Target))
			}
		} else {
			fmt.Println("Button action not recognized in click.Plugin")
		}
	})

	bl.OnMouseMove( func(e *bl.MouseMoveEvent) {
		if lastNodeID == e.Target.Id {
			// we have a click!
			if cb != nil {
				cb(newEvent(e.X, e.Y, e.Target))
			}
		}
	})
}

func NewPlugin() *Plugin {
	plugin = &Plugin{}

	return plugin
}

