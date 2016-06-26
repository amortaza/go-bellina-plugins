package pad

import (
	"github.com/amortaza/go-bellina"
)

var Param_Left = "left"
var Param_Top = "top"
var Param_Right = "right"
var Param_Bottom = "bottom"

type State struct {
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

func (c *Plugin) GetState() interface{} {
	return getOrCreateState(bl.Current_Node.ID)
}

func (c *Plugin) Tick() {
}

func runLogic(shadow *bl.ShadowNode, state *State) {
	node := bl.GetNodeByID(shadow.ID)

	padLeft := bl.GetI_fromNodeID(shadow.ID, "pad", Param_Left)
	padTop := bl.GetI_fromNodeID(shadow.ID, "pad", Param_Top)
	padRight := bl.GetI_fromNodeID(shadow.ID, "pad", Param_Right)
	//padBottom := bl.GetI_fromNodeID(shadow.ID, "pad", Param_Bottom)

	for e := node.Kids.Front(); e != nil; e = e.Next() {
	    	kid := e.Value.(*bl.Node)

		kidShadow := bl.EnsureShadowByID(kid.ID)

		if kidShadow.Left < padLeft {
			kidShadow.Left = padLeft
		}

		if kidShadow.Top < padTop {
			kidShadow.Top = padTop
		}

		right := kidShadow.Left + kidShadow.Width - 1

		if right > shadow.Width - padRight {
			kid.Width = shadow.Width - padRight - kid.Left
		}
	}
}

func (c *Plugin) On(cb func(interface{})) {
	state := getOrCreateState(bl.Current_Node.ID)
	shadow := bl.EnsureShadow()

	for e := bl.Current_Node.Kids.Front(); e != nil; e = e.Next() {
		kid := e.Value.(*bl.Node)
		kidShadow := bl.EnsureShadowByID(kid.ID)

		kid.Left = kidShadow.Left
		kid.Top = kidShadow.Top
		kid.Width = kidShadow.Width
		kid.Height = kidShadow.Height
	}

	bl.AddPluginOnTick( func() {
		runLogic(shadow, state)
	})
}

func SetPaddingAll(pad int32) {
	bl.SetI("pad", Param_Left, pad)
	bl.SetI("pad", Param_Top, pad)
	bl.SetI("pad", Param_Right, pad)
	bl.SetI("pad", Param_Bottom, pad)
}