package horiz

import "github.com/amortaza/go-bellina"

func runLogic(node *bl.Node, state *State) {

	spacing := state.Z_Spacing

	var left = state.Z_Left

	for e := node.Kids.Front(); e != nil; e = e.Next() {

		kid := e.Value.(*bl.Node)

		kid.left = left

		left += kid.width + spacing
	}
}
