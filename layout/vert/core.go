package vert

import (
	"github.com/amortaza/go-bellina"
)

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "vert"
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Init() {
	g_stateByNodeId = make(map[string] *State)
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not supported for vert plugin")
}

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}

