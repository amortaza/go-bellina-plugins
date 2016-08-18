package click

import (
	"github.com/amortaza/go-bellina"
)

var NAME = "click"

var plugin *Plugin

type Plugin struct {
}

func (c *Plugin) Name() string {
	return NAME
}

func (c *Plugin) Init() {
}

func (c *Plugin) Tick() {
}

func (c *Plugin) Reset_ShortTerm() {
}

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
	logic(cb, onDown, onUpAndMiss)
}

