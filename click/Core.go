package click

import (
	"bellina"
	"xel"
	"fmt"
)

func (c *Plugin) On(cb func(interface{})) {

	bl.OnMouseButton( func(e *bl.MouseButtonEvent) {

		if e.Action == xel.Down {
			lastNodeID = e.Target.ID

		} else if e.Action == xel.Up {

			if lastNodeID == e.Target.ID {
				// we have a click!
				cb(Event{bl.Mouse_X, bl.Mouse_X, e.Target})
			}

		} else {
			fmt.Println("Button action not recognized in click.Plugin")
		}
	})
}

