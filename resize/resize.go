package resize

import (
	"github.com/amortaza/go-bellina-plugins/mouse-drag"
	"math"
	"github.com/amortaza/go-bellina"
)

var g_startWidth, g_startHeight int

func Use() {
	On(nil)
}

type Event struct {
	Target *bl.Node
}

func On(cb func(interface{})) {

	shadow := bl.EnsureShadow()

	bl.Dim( shadow.Width, shadow.Height )

	mouse_drag.On_FullLifeCycle(

		// on drag
		func(mouseDragEvent interface{}) {

			e := mouseDragEvent.(mouse_drag.Event)

			diffX := e.MouseX - e.StartX
			diffY := e.MouseY - e.StartY

			width := int(math.Max(float64(g_startWidth + diffX), 16))
			height := int(math.Max(float64(g_startHeight + diffY), 16))

			shadow.Width = width
			shadow.Height = height

			if cb != nil {
				cb(newEvent(e.Target))
			}
		},

		// start drag
		func(mouseDragEvent interface{}) {
			e := mouseDragEvent.(mouse_drag.Event)

			g_startWidth, g_startHeight = e.Target.Width(), e.Target.Height()
		},

		nil)

	shadow.SetDim_on_Node_Only("resize")
}

func newEvent(target *bl.Node) Event {
	return Event{
		target,
	}
}





