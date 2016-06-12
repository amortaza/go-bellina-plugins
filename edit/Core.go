package edit

import (
	"bellina"
	"plugin/focus"
)

var g_editInfoByEditId map[string] *EditInfo

//type Event struct {
//	X, Y int32
//	Target *bl.Node
//}

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "edit"
}

func (c *Plugin) Init() {
	bl.Plugin( focus.NewPlugin() )

	g_editInfoByEditId = make(map[string] *EditInfo)
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not supported for edit plugin")
}

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}

