package vert

import "github.com/amortaza/go-bellina"

var plugin *Plugin

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
	plugin = &Plugin{}

	return plugin
}

func (c *Plugin) GetState() interface{} {
	return ensureState(bl.Current_Node.Id)
}

func (c *Plugin) Tick() {
}

func (c *Plugin) On(cb func(interface{})) {
	state := ensureState(bl.Current_Node.Id)
	shadow := bl.EnsureShadow()

	for e := bl.Current_Node.Kids.Front(); e != nil; e = e.Next() {
		kid := e.Value.(*bl.Node)
		kidShadow := bl.EnsureShadowByID(kid.Id)

		kid.Top = kidShadow.Top
		kid.Height = kidShadow.Height
	}

	bl.AddPluginOnTick(func() {
		runLogic(shadow, state)
	})
}

