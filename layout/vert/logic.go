package vert

import "github.com/amortaza/go-bellina"

func runLogic(shadow *bl.ShadowNode, state *State) {
	node := bl.GetNodeById(shadow.Id)

	spacing := state.Spacing_

	var y = state.Top_
	var kidShadow *bl.ShadowNode
	var pct int

	parentH := shadow.Height

	for e := node.Kids.Front(); e != nil; e = e.Next() {
		kid := e.Value.(*bl.Node)
		kidShadow = bl.EnsureShadowById(kid.Id)

		kidShadow.Top = y

		pct = -1 //bl.GetI_fromNodeID(kid.Id, "vert", "percent")

		if pct > 0 {
			kidShadow.Height = parentH * pct / 100
		}

		y += kidShadow.Height + spacing
	}

	if pct == -1 {
		kidShadow.Height = parentH - kidShadow.Top - 1
	}
}
