package vert

import "github.com/amortaza/go-bellina"

type State struct {
	Z_Top     int
	Z_Spacing int
}

var g_stateById  map[string] *State

func ensureState(nodeId string) *State {
	state, ok := g_stateById[nodeId]

	if !ok {
		state = &State{}

		g_stateById[nodeId] = state
	}

	return state
}

func (s *State) Spacing(spacing int) (*State){
	s.Z_Spacing = spacing

	return s
}

func (s *State) Top(top int) (*State){
	s.Z_Top = top

	return s
}

func (s *State) End() {
	node := bl.Current_Node

	bl.AddStabilizeFunc_PreKids( func() {

		runLogic(node, s)
	});
}

