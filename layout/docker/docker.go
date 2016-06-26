package docker

import (
	"github.com/amortaza/go-bellina"
	"fmt"
)

var Z_ANCHOR_LEFT uint32 = 1 << 0
var Z_ANCHOR_RIGHT uint32 = 1 << 1
var Z_ANCHOR_TOP uint32 = 1 << 2
var Z_ANCHOR_BOTTOM uint32 = 1 << 3

type State struct {
	anchorFlags uint32
}

func (s *State) AnchorBottom() {
	s.anchorFlags |= Z_ANCHOR_BOTTOM
}

func (s *State) AnchorTop() {
	s.anchorFlags |= Z_ANCHOR_TOP
}

func (s *State) AnchorRight() {
	s.anchorFlags |= Z_ANCHOR_RIGHT
}

func (s *State) AnchorLeft() {
	s.anchorFlags |= Z_ANCHOR_LEFT
}

var g_stateByNodeId map[string] *State

func getOrCreateState(nodeId string) *State {
	state, ok := g_stateByNodeId[nodeId]

	if !ok {
		state = &State{}

		g_stateByNodeId[nodeId] = state
	}

	return state
}

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "docker"
}

func (c *Plugin) GetState() interface{} {
	return getOrCreateState(bl.Current_Node.ID)
}

func (c *Plugin) Tick() {
	for key, state := range g_stateByNodeId {
		shadow, _ := bl.GetShadow(key)

		runLogic(shadow, state)
	}
}

func runLogic(shadow *bl.ShadowNode, state *State) {
	getOrCreateState(shadow.ID)
	parentShadow := bl.EnsureShadowByID(shadow.ParentID)

	if state.anchorFlags & Z_ANCHOR_RIGHT != 0 {
		if state.anchorFlags & Z_ANCHOR_LEFT != 0 {
			shadow.Left = 0;
			shadow.Width = parentShadow.Width

		} else {
			shadow.Left = parentShadow.Width - shadow.Width
		}
	} else if state.anchorFlags & Z_ANCHOR_LEFT != 0 {
		shadow.Left = 0;
	}

	if state.anchorFlags & Z_ANCHOR_BOTTOM != 0 {
		if state.anchorFlags & Z_ANCHOR_TOP != 0 {
			shadow.Top = 0;
			shadow.Height = parentShadow.Height

		} else {
			shadow.Top = parentShadow.Height - shadow.Height
		}
	} else if state.anchorFlags & Z_ANCHOR_TOP != 0 {
		shadow.Top = 0;
	}
}

func (c *Plugin) On(cb func(interface{})) {
	getOrCreateState(bl.Current_Node.ID)

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
	fmt.Println("docker init")
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not supported for docker plugin")
}

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}
