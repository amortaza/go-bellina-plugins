package drag_other

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/mouse-drag"
)

type Event struct {
	X, Y int
	Target *bl.Node
}

func Use(nodeId string) {
	On(nodeId, nil)
}

func On(nodeId string, cb func(interface{})) {
	shadow := bl.EnsureShadowById(nodeId)

	shadow.Pos__Node_Only("drag-other")

	cur := bl.Current_Node

	mouse_drag.On( func(mouseDragEvent interface{}) {
		e := mouseDragEvent.(mouse_drag.Event)

		absX, absY := bl.GetNodeAbsolutePos(shadow.BackingNode.Parent)
		shadow.Left = bl.Mouse_X - e.MouseOffsetX - absX - cur.Left
		shadow.Top = bl.Mouse_Y - e.MouseOffsetY - absY - cur.Top

		shadow.Pos__Node_Only("drag-other")

		if cb != nil {
			cb(newEvent(e.Target))
		}
	})
}


