package mouse_drag

import (
	"github.com/amortaza/go-bellina"
	"fmt"
	"github.com/amortaza/go-hal"
)

type Event struct {
	MouseX, MouseY             int
	Target                     *bl.Node
	StartX, StartY             int
	MouseOffsetX, MouseOffsetY int
}

func init() {

	bl.Register_LifeCycle_Before_UserTick_LongTerm(func() {
		gEndCbByNodeId = make(map[string]func(interface{}))
	})

	bl.RegisterLongTerm(bl.EventType_Mouse_Button, func(mouseButtonEvent bl.Event) {

		e := mouseButtonEvent.(*bl.MouseButtonEvent)

		if e.ButtonAction == hal.Button_Action_Up {

			if gLastNodeId != "" {

				endCb, ok := gEndCbByNodeId[gLastNodeId]

				if ok {
					endCb(newEvent(bl.Mouse_X, bl.Mouse_Y, e.Target))
				}
			}

			gLastNodeId = ""
		}
	})

	bl.RegisterLongTerm(bl.EventType_Mouse_Move, func(mouseMoveEvent bl.Event) {

		if gLastNodeId == "" {
			return
		}

		e := mouseMoveEvent.(*bl.MouseMoveEvent)

		node := bl.GetNodeById(gLastNodeId)

		e.Target = node

		node.CallMouseMoveCallbacks(e)
	})

}

func Use() {
	On(nil)
}

func On(cb func(interface{})) {
	On_FullLifeCycle(cb, nil, nil)
}

func On_FullLifeCycle(cb func(interface{}), startCb func(interface{}), endCb func(interface{})) {

	nodeId := bl.Current_Node.Id

	bl.OnMouseButton( func(e *bl.MouseButtonEvent) {
		if e.ButtonAction == hal.Button_Action_Down {

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

		} else if e.ButtonAction == hal.Button_Action_Up {

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


