package pack

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

func runLogic(shadow *bl.ShadowNode, state *State) (width, height int) {

	kids := shadow.BackingNode.Kids

	if state.vertOnly {
		width = shadow.Width
	}

	for kid := kids.Front(); kid != nil; kid = kid.Next() {

		kidNode := kid.Value.(*bl.Node)

		kidShadow := bl.EnsureShadowById(kidNode.Id)

		right := kidShadow.Left + kidShadow.Width
		top := kidShadow.Top + kidShadow.Height

		if !state.vertOnly {

			if right > width {
				width = right
			}
		}

		if top > height {
			height = top
		}
	}

	return width, height
}


