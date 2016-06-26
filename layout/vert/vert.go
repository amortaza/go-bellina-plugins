package vert

import (
	"github.com/amortaza/go-bellina"
)

var Param_Percent = "percent"

type State struct {
	//anchorFlags uint32
}

//func (s *State) AnchorBottom() {
//	s.anchorFlags |= Z_ANCHOR_BOTTOM
//}

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
	//for key, state := range g_stateByNodeId {
	//	shadow, _ := bl.GetShadow(key)
	//
	//	runLogic(shadow, state)
	//}
}

func runLogic(shadow *bl.ShadowNode, state *State) {
	node := bl.GetNodeByID(shadow.ID)

	var x int32 = 0
	var kidShadow *bl.ShadowNode
	var pct int32

	parentH := shadow.Height

	for e := node.Kids.Front(); e != nil; e = e.Next() {
	    	kid := e.Value.(*bl.Node)
		kidShadow = bl.EnsureShadowByID(kid.ID)

		kidShadow.Top = x

		pct = bl.GetI_fromNodeID(kid.ID, "vert", "percent")

		if pct > 0 {
			kidShadow.Height = parentH * pct / 100
		}

		x += kidShadow.Height
	}

	if pct == -1 {
		kidShadow.Height = parentH - kidShadow.Top - 1
	}
}

func (c *Plugin) On(cb func(interface{})) {
	state := getOrCreateState(bl.Current_Node.ID)
	shadow := bl.EnsureShadow()

	for e := bl.Current_Node.Kids.Front(); e != nil; e = e.Next() {
		kid := e.Value.(*bl.Node)
		kidShadow := bl.EnsureShadowByID(kid.ID)

		kid.Top = kidShadow.Top
		kid.Height = kidShadow.Height
	}

	bl.AddPluginOnTick(func() {
		runLogic(shadow, state)
	})
}

