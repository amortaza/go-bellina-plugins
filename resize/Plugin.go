package resize

import (
	"github.com/amortaza/go-bellina-plugins/mouse-drag"
	"math"
	"github.com/amortaza/go-bellina"
	"fmt"
)

var plugin *Plugin

type Event struct {
	Target *bl.Node
}

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "resize"
}

func (c *Plugin) GetState() interface{} {
	return nil
}

func (c *Plugin) Tick() {
}

func (c *Plugin) Reset() {
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Init() {
}

func (c *Plugin) Uninit() {
}

var startWidth, startHeight int

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not supported for Resiable plugin")
}

func (c *Plugin) On(cb func(interface{})) {

	shadow := bl.EnsureShadow()

	bl.Dim( shadow.Width, shadow.Height )

	mouse_drag.On2(

		// on drag
		func(mouseDragEvent interface{}) {


			e := mouseDragEvent.(mouse_drag.Event)
			fmt.Println("drag ", e.Target.Id)

			shadow := bl.EnsureShadowById(e.Target.Id)

			diffX := e.MouseX - e.StartX
			diffY := e.MouseY - e.StartY

			width := int(math.Max(float64(startWidth + diffX), 16))
			height := int(math.Max(float64(startHeight + diffY), 16))

			shadow.Dim(width, height)

			if cb != nil {
				cb(newEvent(e.Target))
			}
		},

		// start drag
		func(mouseDragEvent interface{}) {
			e := mouseDragEvent.(mouse_drag.Event)

			startWidth, startHeight = e.Target.Width, e.Target.Height
		},

		nil)

	bl.EnsureShadow().Dim_to_Node()
	
	//fmt.Println(bl.Current_Node.Id, bl.Current_Node.Width)
}

func newEvent(target *bl.Node) Event {
	return Event{
		target,
	}
}

