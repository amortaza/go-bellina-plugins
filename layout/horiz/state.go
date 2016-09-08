package horiz

import "github.com/amortaza/go-bellina"

type State struct {
	Z_Left    int
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

func (s *State) Left(left int) (*State){
	s.Z_Left = left

	return s
}

func (s *State) End() {
	node := bl.Current_Node

	bl.AddFunc(func() {
		runLogic(node, s)
	})
}

