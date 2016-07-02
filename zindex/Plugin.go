package zindex

import (
	"github.com/amortaza/go-bellina"
	"fmt"
)

var plugin *Plugin


type Plugin struct {
}

func (c *Plugin) Name() string {
	return "zindex"
}

func (c *Plugin) Init() {
	g_ctxByNodeId = make(map[string] *Ctx)
}

func (c *Plugin) Tick() {
}

func (c *Plugin) Reset() {
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
	delete(g_ctxByNodeId, node.Id)
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	fmt.Println("On2 not supoorted in zindex plugin")
}


