package vert

import "github.com/amortaza/go-bellina"

func runLogic(node *bl.Node, state *State) {

	spacing := state.spacing

	var y = state.top

	for e := node.Kids.Front(); e != nil; e = e.Next() {

		kid := e.Value.(*bl.Node)

		kidShadow := bl.EnsureShadowByNode(kid)

		kidShadow.Top = y

		y += kidShadow.Height + spacing
	}
}
