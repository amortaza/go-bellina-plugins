package horiz

import "github.com/amortaza/go-bellina"

func runLogic(node *bl.Node, state *State) {

	spacing := state.spacing

	var left = state.left

	for e := node.Kids.Front(); e != nil; e = e.Next() {

		kid := e.Value.(*bl.Node)

		kid.SetLeft(left)

		left += kid.Width() + spacing
	}
}
