package vert

type State struct {
	VertId          string

	Top_     	int
	Spacing_	int
	Percent_	int
}

var g_stateById  map[string] *State

func ensureState(vertId string) *State {
	state, ok := g_stateById[vertId]

	if !ok {
		state = &State{VertId: vertId}

		g_stateById[vertId] = state
	}

	return state
}

func (s *State) Spacing(spacing int) (*State){
	s.Spacing_ = spacing

	return s
}

func (s *State) Top(top int) (*State){
	s.Top_ = top

	return s
}

func (s *State) Percent(percent int) (*State){
	s.Percent_ = percent

	return s
}

func (s *State) End() (*State){
	Use()

	return s
}

