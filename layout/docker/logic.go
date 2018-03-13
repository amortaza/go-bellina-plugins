package docker

import (
	"github.com/amortaza/go-bellina"
)

var g_stateByNodeId map[string] *State

func ensureState(nodeId string) *State {
	state, ok := g_stateByNodeId[nodeId]

	if !ok {
		state = &State{ sudo: "docker"}

		g_stateByNodeId[nodeId] = state
	}

	return state
}

func runLogic(node *bl.Node, state *State) (left, top, width, height int) {

	parentNode := node.Parent

	left, top, width, height = node.Left(), node.Top(), node.Width(), node.Height()

	// right
	if state.anchorFlags & _ANCHOR_RIGHT != 0 {

		// left AND right
		if state.anchorFlags & _ANCHOR_LEFT != 0 {

			left = state.leftPadding;
			width = parentNode.Width() - state.leftPadding - state.rightPadding

			if width < 16 {
				width = 16
			}

		} else {
			// right only
			left = parentNode.Width() - node.Width() - state.rightPadding
		}
	} else if state.anchorFlags & _ANCHOR_LEFT != 0 {
		// left only
		left = state.leftPadding;
	}

	// bottom
	if state.anchorFlags & _ANCHOR_BOTTOM != 0 {

		// bottom AND top
		if state.anchorFlags & _ANCHOR_TOP != 0 {
			top = state.topPadding;
			height = parentNode.Height() - state.topPadding - state.bottomPadding

		} else {
			// bottom only
			top = parentNode.Height() - node.Height() - state.bottomPadding
		}
	} else if state.anchorFlags & _ANCHOR_TOP != 0 {
		// top only
		top = state.topPadding;
	}

	return left, top, width, height
}


