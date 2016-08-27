package vert

import "github.com/amortaza/go-bellina"

func runLogic(shadow *bl.ShadowNode, state *State) {
	node := bl.GetNodeById(shadow.Id)

	spacing := state.Spacing_

	var y = state.Top_
	var kidShadow *bl.ShadowNode

	for e := node.Kids.Front(); e != nil; e = e.Next() {
		kid := e.Value.(*bl.Node)
		kidShadow = bl.EnsureShadowById(kid.Id)

		kidShadow.Top = y

		y += kidShadow.Height + spacing
	}
}
