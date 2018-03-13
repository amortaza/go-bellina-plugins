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

	//shadow.SetPos_on_Node_Only("drag")
	shadow.SetLeft_on_Node_Only("drag")
	shadow.SetTop_on_Node_Only("drag")

	mouse_drag.On( func(mouseDragEvent interface{}) {

		e := mouseDragEvent.(mouse_drag.Event)

		absX, absY := bl.GetNodeAbsolutePos(e.Target.Parent)
		shadow.Left = bl.Mouse_X - e.MouseOffsetX - absX
		shadow.Top = bl.Mouse_Y - e.MouseOffsetY - absY

		if cb != nil {
			cb(newEvent(e.Target))
		}
	})
}


