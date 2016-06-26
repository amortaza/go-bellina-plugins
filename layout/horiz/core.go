package horiz

import (
	"github.com/amortaza/go-bellina"
	"fmt"
)

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "horiz"
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Init() {
	g_stateByNodeId = make(map[string] *State)
	fmt.Println("horiz init")
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not supported for horiz plugin")
}

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}

