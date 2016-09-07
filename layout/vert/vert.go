package vert

import (
	"github.com/amortaza/go-bellina"
)

// Shared variable across Div()/End()
var g_curState *State

func init() {
	g_stateById = make(map[string] *State)
}

func Id() *State {
	g_curState = ensureState(bl.Current_Node.Id)

	return g_curState
}

