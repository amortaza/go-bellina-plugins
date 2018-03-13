package stretch_stick

import "github.com/amortaza/go-bellina"

func StickTo(otherId string) *State {

	g_curState = ensureState(bl.Current_Node.Id)

	g_curState.otherId = otherId

	return g_curState
}

func End() {

	other := bl.GetNodeById(g_curState.otherId)

	other.SetOwnerOfWidth("stretch-stick")

	//otherRight :=
}
