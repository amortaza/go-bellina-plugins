package vert

import (
	"github.com/amortaza/go-bellina"
)

var g_state *State

func init() {

	g_stateById = make(map[string] *State)
}

func Use() *State {

	g_state = ensureState(bl.Current_Node.Id)

	return g_state
}

