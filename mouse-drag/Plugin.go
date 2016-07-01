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
	gEndCbByNodeId = make(map[string] func(interface{}))
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Init() {
	gEndCbByNodeId = make(map[string] func(interface{}))

	event.RegisterLongTerm(bl.EventType_Mouse_Button, func(mouseButtonEvent event.Event) {

		e := mouseButtonEvent.(*bl.MouseButtonEvent)

		if e.ButtonAction == xel.Button_Action_Up {
			if gLastNodeId != "" {
				endCb, ok := gEndCbByNodeId[gLastNodeId]

				if ok {
					endCb(newEvent(bl.Mouse_X, bl.Mouse_Y, e.Target))
				}
			}

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

	nodeId := bl.Current_Node.Id

	bl.OnMouseButton( func(e *bl.MouseButtonEvent) {
		if e.ButtonAction == xel.Button_Action_Down {

			// target == node when target is child of node!!
			if e.Target.Id != nodeId {
				return
			}

			gLastNodeId = e.Target.Id

			gStartX, gStartY = bl.Mouse_X, bl.Mouse_Y

			absX, absY := bl.GetNodeAbsolutePos(e.Target)
			gMouseOffsetX = bl.Mouse_X - absX
			gMouseOffsetY = bl.Mouse_Y - absY

			if startCb != nil {
				startCb(newEvent(bl.Mouse_X, bl.Mouse_Y, e.Target))
			}

		} else if e.ButtonAction == xel.Button_Action_Up {

			// lastnodeid CAN be empty
			if endCb != nil && gLastNodeId == nodeId {
				endCb(newEvent(bl.Mouse_X, bl.Mouse_Y, e.Target))
			}

			gLastNodeId = ""

		} else {
			fmt.Println("Button action not recognized in click.Plugin")
		}
	})

	bl.OnMouseMove( func(e *bl.MouseMoveEvent) {
		if gLastNodeId != e.Target.Id {
			return
		}

		// we have a drag!
		if cb != nil {
			cb(newEvent(e.X, e.Y, e.Target))
		}
	})

	if endCb != nil {
		gEndCbByNodeId[nodeId] = endCb
	}
}


