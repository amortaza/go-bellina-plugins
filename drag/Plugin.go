package drag

import (
	"bellina"
	"plugin/mouse-drag"
	"fmt"
)

type Event struct {
	X, Y int32
	Target *bl.Node
}

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "drag"
}

func (c *Plugin) Init() {
	bl.Plugin( mouse_drag.NewPlugin() )
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	fmt.Println("On2 not supported for draggable plugin")
}

func (c *Plugin) On(cb func(interface{})) {

	shadow := bl.EnsureShadow()

	bl.Pos( shadow.Left, shadow.Top )

	bl.On("mouse-drag", func(mouseDragEvent interface{}) {
		e := mouseDragEvent.(mouse_drag.Event)

		shadow, _ := bl.GetShadow(e.Target.ID)

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

func newEvent(target *bl.Node) Event {
	return Event{
		bl.Mouse_X, bl.Mouse_X,
		target,
	}
}

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}
