package stretch_stick

import "github.com/amortaza/go-bellina"

type State struct {

	otherId string

	anchorLeft bool
}

var g_stateById map[string] *State
var g_curState *State

func (state *State) AnchorLeft() *State {

	state.anchorLeft = true

	return state
}

func (state *State) End() {

	End()
}

func init() {

	g_stateById = make(map[string] *State)
}

func ensureState(nodeId string) *State {

	state, ok := g_stateById[nodeId]

	if !ok {

		state = &State{}

		g_stateById[nodeId] = state

		bl.OnFreeNode(onFreeNode)
	}

	return state
}

func onFreeNode(nodeId string) {

	delete(g_stateById, nodeId)
}
