package drag

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/mouse-drag"
)

type Event struct {
	X, Y int
	Target *bl.Node
}

func Use() {
	On(nil)
}

func On(cb func(interface{})) {

	shadow := bl.EnsureShadow()

	shadow.SetLeft_on_Node_Only("drag")
	shadow.SetTop_on_Node_Only("drag")

	mouse_drag.On( func(mouseDragEvent interface{}) {

		e := mouseDragEvent.(mouse_drag.Event)

		absX, absY := bl.GetNodeAbsolutePos(e.Target.Parent)

		x := bl.Mouse_X - e.MouseOffsetX - absX
		y := bl.Mouse_Y - e.MouseOffsetY - absY

		state, ok := g_stateById[ e.Target.Id]

		var pipeTo func(x, y int)

		if ok {

			pipeTo = state.pipeTo
		}

		if pipeTo != nil {
			pipeTo(x, y)

		} else {

			shadow.Left = x
			shadow.Top = y

			if cb != nil {
			cb(newEvent(e.Target))
			}
		}
	})
}

func PipeTo(pipeTo func(x, y int)) {

	state := ensureState(bl.Current_Node.Id)

	state.pipeTo = pipeTo
}

