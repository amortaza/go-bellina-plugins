package click

import (
	"github.com/amortaza/go-bellina/event"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-xel"
	"fmt"
)

var gLastNodeId string

func logic(cb func(interface{}), onDown func(interface{}), onUpAndMiss func(interface{})) {

	nodeId := bl.Current_Node.Id

	event.RegisterShortTerm(bl.EventType_Mouse_Button, func(event event.Event) {

		e := event.(*bl.MouseButtonEvent)

		if e.ButtonAction == xel.Button_Action_Down {

			if e.Target.Id != nodeId {
				return
			}

			gLastNodeId = e.Target.Id

			if onDown != nil {
				onDown(Event{bl.Mouse_X, bl.Mouse_X, e.Target})
			}

		} else if e.ButtonAction == xel.Button_Action_Up {

			if gLastNodeId == e.Target.Id {
				// we have a click!

				if cb != nil {
					cb(Event{bl.Mouse_X, bl.Mouse_X, e.Target})
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


