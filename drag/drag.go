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

	bl.Pos( shadow.Left, shadow.Top )

	mouse_drag.On( func(mouseDragEvent interface{}) {
		e := mouseDragEvent.(mouse_drag.Event)

		shadow := bl.EnsureShadowById(e.Target.Id)

		absX, absY := bl.GetNodeAbsolutePos(e.Target.Parent)
		shadow.Left = bl.Mouse_X - e.MouseOffsetX - absX
		shadow.Top = bl.Mouse_Y - e.MouseOffsetY - absY

		e.Target.Left = shadow.Left
		e.Target.Top = shadow.Top

		if cb != nil {
			cb(newEvent(e.Target))
		}
	})
}


