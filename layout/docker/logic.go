package docker

import "github.com/amortaza/go-bellina"

var g_stateByNodeId map[string] *State

func ensureState(nodeId string) *State {
	state, ok := g_stateByNodeId[nodeId]

	if !ok {
		state = &State{}

		g_stateByNodeId[nodeId] = state
	}

	return state
}

func runLogic(shadow *bl.ShadowNode, state *State) {
	ensureState(shadow.Id)
	parentShadow := bl.EnsureShadowByID(shadow.ParentID)

	if state.anchorFlags & Z_ANCHOR_RIGHT != 0 {
		if state.anchorFlags & Z_ANCHOR_LEFT != 0 {
			shadow.Left = 0;
			shadow.Width = parentShadow.Width

		} else {
			shadow.Left = parentShadow.Width - shadow.Width
		}
	} else if state.anchorFlags & Z_ANCHOR_LEFT != 0 {
		shadow.Left = 0;
	}

	if state.anchorFlags & Z_ANCHOR_BOTTOM != 0 {
		if state.anchorFlags & Z_ANCHOR_TOP != 0 {
			shadow.Top = 0;
			shadow.Height = parentShadow.Height

		} else {
			shadow.Top = parentShadow.Height - shadow.Height
		}
	} else if state.anchorFlags & Z_ANCHOR_TOP != 0 {
		shadow.Top = 0;
	}
}


