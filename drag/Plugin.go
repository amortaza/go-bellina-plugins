package drag

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/mouse-drag"
	"fmt"
)

var plugin *Plugin

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "drag"
}

func (c *Plugin) Init() {
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Tick() {
}

func (c *Plugin) Reset_ShortTerm() {
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	fmt.Println("On2 not supported for drag plugin")
}

func (c *Plugin) On(cb func(interface{})) {

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


