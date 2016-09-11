package padsize

import "github.com/amortaza/go-bellina"

type State struct {
	Z_Left    int
	Z_Top    int
	Z_Right    int
	Z_Bottom    int
}

var g_stateById  map[string] *State

func init() {
	g_stateById = make(map[string] *State)
}

func ensureState(nodeId string) *State {
	state, ok := g_stateById[nodeId]

	if !ok {
		state = &State{}

		g_stateById[nodeId] = state
	}

	return state
}

func (state *State) Left(value int) (*State){
	state.Z_Left = value

	return state
}

func (state *State) Top(value int) (*State){
	state.Z_Top = value

	return state
}

func (state *State) Right(value int) (*State){
	state.Z_Right = value

	return state
}

func (state *State) Bottom(value int) (*State){
	state.Z_Bottom = value

	return state
}

func (state *State) End() {
	node := bl.Current_Node

	bl.AddFunc( func() {
		runLogic(node, state)
	})
}

