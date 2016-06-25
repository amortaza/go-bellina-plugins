package focus

import (
	"github.com/amortaza/go-bellina"
)

func newEvent(target *bl.Node, keyEvent *bl.KeyEvent) Event {
	return Event{target, keyEvent}
}

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) Tick() {
}

