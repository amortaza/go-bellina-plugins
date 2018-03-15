package drag

import "github.com/amortaza/go-bellina"

func newEvent(target *bl.Node) Event {

	return Event{

		bl.Mouse_X, bl.Mouse_X,
		target,
	}
}

type State struct {

	pipeTo func(x, y int)
}

var g_stateById map[string] *State

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

