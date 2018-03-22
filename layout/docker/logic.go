package docker

import (
	"github.com/amortaza/go-bellina"
)

var g_stateByNodeId map[string] *State

func ensureState(nodeId string) *State {
	state, ok := g_stateByNodeId[nodeId]

	if !ok {
		state = &State{}

		g_stateByNodeId[nodeId] = state
	}

	return state
}

func runLogic(shadow *bl.ShadowNode, state *State) (left, top, width, height int) {

	parentShadow := bl.EnsureShadowById(shadow.BackingNode.Parent.Id)

	left, top, width, height = shadow.Left, shadow.Top, shadow.Width, shadow.Height

	// right
	if state.anchorFlags & _ANCHOR_RIGHT != 0 {

		// left AND right
		if state.anchorFlags & _ANCHOR_LEFT != 0 {

			left = state.leftPadding;
			width = parentShadow.Width - state.leftPadding - state.rightPadding

			if width < 16 {
				width = 16
			}

		} else {
			// right only
			left = parentShadow.Width - shadow.Width - state.rightPadding
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
			height = parentShadow.Height - state.topPadding - state.bottomPadding

		} else {
			// bottom only
			top = parentShadow.Height - shadow.Height - state.bottomPadding
		}
	} else if state.anchorFlags & _ANCHOR_TOP != 0 {
		// top only
		top = state.topPadding;
	}

	return left, top, width, height
}


