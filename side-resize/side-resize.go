package side_resize

import (
	"github.com/amortaza/go-bellina-plugins/mouse-drag"
	"github.com/amortaza/go-bellina"
	"fmt"
)

func fake() {
	var _ = fmt.Print
}

func NodeId(otherId string) (*State){
	g_otherId = otherId
	g_flags = 0
	g_sudo = "side-resize"

    	return &State{}
}

func logic() {
	validate()

    	shadowOther := bl.EnsureShadowById(g_otherId)
	curShadow := bl.EnsureShadow()

	flags := g_flags
	cur := bl.Current_Node

    	set(shadowOther, curShadow)

    	mouse_drag.On_FullLifeCycle(

        	// on drag
		func(mouseDragEvent interface{}) {

			e := mouseDragEvent.(mouse_drag.Event)

			mouseDiffX := e.MouseX - e.StartX
			mouseDiffY := e.MouseY - e.StartY

			absX_otherParent, absY_otherParent := bl.GetNodeAbsolutePos(shadowOther.BackingNode.Parent)
			absX_curParent, absY_curParent := bl.GetNodeAbsolutePos(curShadow.BackingNode.Parent)

			var _ = absX_otherParent
			var _ = absY_otherParent
			//var _ = absY_curParent

			if flags & _LEFT != 0 {
				otherRight := shadowOther.Left + shadowOther.Width

				shadowOther.Left = bl.Mouse_X - e.MouseOffsetX - cur.Left
				shadowOther.Width = otherRight - shadowOther.Left
			}

			if flags & _RIGHT != 0 {
				curShadow.Left = bl.Mouse_X - e.MouseOffsetX - absX_curParent
				shadowOther.Width = g_startWidth + mouseDiffX
			}

			if flags & _TOP != 0 {
				otherBottom := shadowOther.Top + shadowOther.Height

				shadowOther.Top = bl.Mouse_Y - e.MouseOffsetY - cur.Top
				shadowOther.Height = otherBottom - shadowOther.Top
			}

			if flags & _BOTTOM != 0 {
				curShadow.Top = bl.Mouse_Y - e.MouseOffsetY - absY_curParent
				shadowOther.Height = g_startHeight + mouseDiffY
			}

			set(shadowOther, curShadow)

			//if cb != nil {
			//  cb(newEvent(e.Target))
			//}
		},

		// start drag
		func(mouseDragEvent interface{}) {
			g_startWidth, g_startHeight = shadowOther.BackingNode.Width, shadowOther.BackingNode.Height
		},

		// end drag??
		nil)
}

