package docker

import (
	"github.com/amortaza/go-bellina"
)

var plugin *Plugin

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "docker"
}

func (c *Plugin) GetState() interface{} {
	return ensureState(bl.Current_Node.Id)
}

func (c *Plugin) Tick() {
	for key, state := range g_stateByNodeId {
		shadow, _ := bl.GetShadowById(key)

		runLogic(shadow, state)
	}
}

func (c *Plugin) On(cb func(interface{})) {
	ensureState(bl.Current_Node.Id)

	shadow := bl.EnsureShadow()

	bl.Current_Node.Left = shadow.Left
	bl.Current_Node.Top = shadow.Top
	bl.Current_Node.Width = shadow.Width
	bl.Current_Node.Height = shadow.Height
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
	panic("On2 not supported for docker plugin")
}

func NewPlugin() *Plugin {
	plugin = &Plugin{}

	return plugin
}

