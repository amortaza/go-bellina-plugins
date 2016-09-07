package docker

import (
	"github.com/amortaza/go-bellina"
	"fmt"
)

func fake() {
    var _ = fmt.Println
}

var g_stateByNodeId map[string] *State

func ensureState(nodeId string) *State {
	state, ok := g_stateByNodeId[nodeId]

	if !ok {
		state = &State{}

		g_stateByNodeId[nodeId] = state
	}

	return state
}

func runLogic(node *bl.Node, state *State) (left, top, width, height int) {
	parentNode := node.Parent

	left, top, width, height = node.Left, node.Top, node.Width, node.Height

	// to do
	//fmt.Println("(1 docker.runLogic) Parent Width ", parentNode.Id, " : ", parentNode.Width)

	if state.anchorFlags & _ANCHOR_RIGHT != 0 {
		if state.anchorFlags & _ANCHOR_LEFT != 0 {
			
			//fmt.Println("here", )
			//bl.Disp(parentNode)

			left = 0;
			width = parentNode.Width

		} else {
			left = parentNode.Width - node.Width
		}
	} else if state.anchorFlags & _ANCHOR_LEFT != 0 {
		left = 0;
	}

	if state.anchorFlags & _ANCHOR_BOTTOM != 0 {
		if state.anchorFlags & _ANCHOR_TOP != 0 {
			top = 0;
			height = parentNode.Height

		} else {
			top = parentNode.Height - node.Height
		}
	} else if state.anchorFlags & _ANCHOR_TOP != 0 {
		top = 0;
	}

	return left, top, width, height
}


