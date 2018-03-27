package pack

import (
	"github.com/amortaza/go-bellina"
)

func init() {

	g_stateByNodeId = make(map[string] *State)
}

type State struct {

	vertOnly bool
}

func Use() (state *State){

	node := bl.Current_Node

	shadow := bl.EnsureShadow()

	state = ensureState(node.Id)

	bl.AddStabilizeFunc_PostKids(func() {

		shadow.Width, shadow.Height = runLogic(shadow, state)
	})

	return state
}

func (state *State) Vert() (*State) {

	state.vertOnly = true

	return state
}

func (state *State) End() {

	shadow := bl.EnsureShadow()

	bl.AddStabilizeFunc_PostKids(func() {

		shadow.Width, shadow.Height = runLogic(shadow, state)
	})
}


