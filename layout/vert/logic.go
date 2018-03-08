package vert

import "github.com/amortaza/go-bellina"

func runLogic(node *bl.Node, state *State) {

	spacing := state.Z_Spacing

	var y = state.Z_Top

	for e := node.Kids.Front(); e != nil; e = e.Next() {

		kid := e.Value.(*bl.Node)

		kid.top = y

		y += kid.height + spacing
	}
}
