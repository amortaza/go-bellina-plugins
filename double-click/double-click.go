package double_click

import (
	"github.com/amortaza/go-bellina-plugins/click"
	"time"
	"github.com/amortaza/go-bellina"
)

var gLastNodeID string
var gLastMs int64
var gSpeedMs int64

func init() {
	gSpeedMs = 1000
}

type Event struct {
	X, Y int
	Target *bl.Node
}

func On(cb func(interface{})) {

	click.On(func(i interface{}) {
		e := i.(click.Event)

		if gLastNodeID == "" {
			gLastNodeID = e.Target.Id
			gLastMs = time.Now().UnixNano() / 1e6

		} else if gLastNodeID == e.Target.Id {
			nowMs := time.Now().UnixNano() / 1e6

			if nowMs - gLastMs < gSpeedMs {
				// we have a double-click!
				cb(Event{bl.Mouse_X, bl.Mouse_X, e.Target})
				gLastNodeID = ""

			} else {
				gLastNodeID = ""
			}

		} else {
			gLastNodeID = ""
		}
	})
}

