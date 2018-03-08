package resize_other

import (
	"github.com/amortaza/go-bellina-plugins/mouse-drag"
	"math"
	"github.com/amortaza/go-bellina"
)

var g_startWidth, g_startHeight int

func Use(nodeId string) {
	On(nodeId, nil)
}

type Event struct {
	Target *bl.Node
}

func On(nodeId string, cb func(interface{})) {

	shadow := bl.EnsureShadowById(nodeId)

	shadow.Dim__Node_Only("resize-other")

	mouse_drag.On_FullLifeCycle(

		// on drag
		func(mouseDragEvent interface{}) {

			e := mouseDragEvent.(mouse_drag.Event)

			diffX := e.MouseX - e.StartX
			diffY := e.MouseY - e.StartY

			width := int(math.Max(float64(g_startWidth + diffX), 16))
			height := int(math.Max(float64(g_startHeight + diffY), 16))

			shadow.Dim__Self_and_Node(width, height, "resize-other")

			if cb != nil {
				cb(newEvent(e.Target))
			}
		},

		// start drag
		func(mouseDragEvent interface{}) {

			g_startWidth, g_startHeight = shadow.BackingNode.width, shadow.BackingNode.height
		},

		nil)

	shadow.Dim__Node_Only("resize-other")
}

func newEvent(target *bl.Node) Event {
	return Event{
		target,
	}
}





