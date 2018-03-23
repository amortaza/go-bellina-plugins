package horiz_split_pane

import "github.com/amortaza/go-bellina"

type State struct {

	leftId, handleId, rightId string
}

var g_stateById map[string] *State

func init() {

	g_stateById = make(map[string] *State)
}

func ensureState() (*State, bool) {

	splitParentId := bl.Current_Node.Parent.Id

	state, ok := g_stateById[splitParentId]

	if !ok {

		state = &State{}

		g_stateById[splitParentId] = state

		bl.OnFreeNodeId(splitParentId, onFreeNode)
	}

	return state, !ok
}

func onFreeNode(nodeId string) {

	delete(g_stateById, nodeId)
}

