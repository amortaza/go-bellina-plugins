package horiz

import "github.com/amortaza/go-bellina"

func runLogic(node *bl.Node, state *State) {

	spacing := state.spacing

	var x = state.left

	for e := node.Kids.Front(); e != nil; e = e.Next() {

		kid := e.Value.(*bl.Node)

		kidShadow := bl.EnsureShadowByNode(kid)

		kidShadow.Left = x

		x += kidShadow.Width + spacing
	}
}
