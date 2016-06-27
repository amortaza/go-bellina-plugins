package focus

import (
	"github.com/amortaza/go-bellina"
)

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) Tick() {
}

func (e *Event) Type() string {
	return "focus"
}

func (p *Plugin) Name() string {
	return "focus"
}

func (c *Plugin) GetState() interface{} {
	return nil
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (p *Plugin) On(cb func(interface{})) {
	p.On2(cb, nil, nil)
}
