package click

import (
	"github.com/amortaza/go-bellina"
	"fmt"
)

var gLastNodeId string

func logic(cb, onDown, onUpAndMiss func(interface{})) {

	nodeId := bl.Current_Node.Id

	bl.RegisterShortTerm(bl.EventType_Mouse_Button, func(event bl.Event) {

		e := event.(*bl.MouseButtonEvent)

		if e.ButtonAction == bl.Button_Action_Down {

			if e.Target.Id != nodeId {
				return
			}

			gLastNodeId = e.Target.Id

			if onDown != nil {
				onDown(Event{bl.Mouse_X, bl.Mouse_X, e.Target})
			}

		} else if e.ButtonAction == bl.Button_Action_Up {

			if gLastNodeId == e.Target.Id {
				// we have a click!

				if cb != nil && gLastNodeId == nodeId {
					cb(Event{e.X, e.Y, e.Target})
				}

			} else if gLastNodeId == nodeId {
				gLastNodeId = ""

				if onUpAndMiss != nil {
					onUpAndMiss(nil)
				}
			}

		} else {
			fmt.Println("Button action not recognized in click.Plugin")
		}
	})
}


