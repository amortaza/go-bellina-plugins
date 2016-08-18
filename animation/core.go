package animation

import (
	"github.com/amortaza/go-bellina"
	"container/list"
)

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On(cb func(interface{})) {
	c.On2(cb, nil, nil)
}

func (c *Plugin) On2(cb func(interface{}), onDown func(interface{}), onUpAndMiss func(interface{})) {
}

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "animation"
}

func (c *Plugin) Init() {
	g_states = list.New()
}

func (c *Plugin) Reset_ShortTerm() {
}


