package resize

import (
	"github.com/amortaza/go-bellina-plugins/mouse-drag"
	"math"
	"github.com/amortaza/go-bellina"
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

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Init() {
	bl.Plugin( mouse_drag.NewPlugin() )
}

func (c *Plugin) Uninit() {
}

var startWidth, startHeight int32

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

			shadow, _ := bl.GetShadowById(e.Target.Id)

			diffX := e.X - e.StartX
			diffY := e.Y - e.StartY

			shadow.Width = int32(math.Max(float64(startWidth + diffX), 16))
			shadow.Height = int32(math.Max(float64(startHeight + diffY), 16))

			if cb != nil {
				e.Target.Width, e.Target.Height = shadow.Width, shadow.Height

				cb(newEvent(e.Target))
			}
		},

		// start drag
		func(mouseDragEvent interface{}) {
			e := mouseDragEvent.(mouse_drag.Event)

			startWidth, startHeight = e.Target.Width, e.Target.Height
		},

		nil)
}

func newEvent(target *bl.Node) Event {
	return Event{
		target,
	}
}

func NewPlugin() *Plugin {
	plugin = &Plugin{}

	return plugin
}
