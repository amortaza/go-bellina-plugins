package horiz

import "github.com/amortaza/go-bellina"

type State struct {

	left    int
	spacing int
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

	s.spacing = spacing

	return s
}

func (s *State) Left(left int) (*State){

	s.left = left

	return s
}

func (s *State) End() {

	node := bl.Current_Node

	bl.AddStabilizeFunc_PreKids(func() {

		runLogic(node, s)
	})
}

