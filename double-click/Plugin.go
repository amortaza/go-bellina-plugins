package double_click

import (
	"github.com/amortaza/go-bellina-plugins/click"
	"time"
	"github.com/amortaza/go-bellina"
)

var lastNodeID string
var lastMs int64

type Event struct {
	X, Y int32
	Target *bl.Node
}

type Plugin struct {
	speedMs int64
}

func (c *Plugin) Name() string {
	return "double-click"
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
	bl.Plugin( click.NewPlugin() )
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not supported for double-click plugin")
}

func (c *Plugin) On(cb func(interface{})) {

	click.On(func(i interface{}) {
		e := i.(click.Event)

		if lastNodeID == "" {
			lastNodeID = e.Target.Id
			lastMs = time.Now().UnixNano() / 1e6

		} else if lastNodeID == e.Target.Id {
			nowMs := time.Now().UnixNano() / 1e6

			if nowMs - lastMs < c.speedMs {
				// we have a double-click!
				cb(Event{bl.Mouse_X, bl.Mouse_X, e.Target})
				lastNodeID = ""

			} else {
				lastNodeID = ""
			}

		} else {
			lastNodeID = ""
		}
	})
}

func NewPlugin(speedMs int64) *Plugin {
	if speedMs == 0 {
		speedMs = 1000
	}

	c := &Plugin{speedMs}

	return c
}
