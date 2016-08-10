package vert

import (
	//"github.com/amortaza/go-bellina"
)
/*
var g_stateByNodeId map[string] *State

func ensureState(nodeId string) *State {
	state, ok := g_stateByNodeId[nodeId]

	if !ok {
		state = &State{}

		g_stateByNodeId[nodeId] = state
	}

	return state
}

func runLogic(shadow *bl.ShadowNode, state *State) {
	node := bl.GetNodeByID(shadow.Id)

	spacing := bl.GetI_fromNodeID( shadow.Id, "vert", "spacing" )

	var y = 0
	var kidShadow *bl.ShadowNode
	var pct int

	parentH := shadow.Height

	for e := node.Kids.Front(); e != nil; e = e.Next() {
		kid := e.Value.(*bl.Node)
		kidShadow = bl.EnsureShadowByID(kid.Id)

		kidShadow.Top = y

		pct = bl.GetI_fromNodeID(kid.Id, "vert", "percent")

		if pct > 0 {
			kidShadow.Height = parentH * pct / 100
		}

		y += kidShadow.Height + spacing
	}

	if pct == -1 {
		kidShadow.Height = parentH - kidShadow.Top - 1
	}
}
*/