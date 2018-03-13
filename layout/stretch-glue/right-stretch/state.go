package stretch_glue_right

import "github.com/amortaza/go-bellina"

type State struct {
	Z_Left_NodeId	string
	Z_Right_NodeId	string
}

func ensureState() *State {
	return &State{}
}

func (s *State) LeftNodeId(nodeId string) (*State){
	s.Z_Left_NodeId = nodeId

	return s
}

func (s *State) RightNodeId(nodeId string) (*State){
	s.Z_Right_NodeId = nodeId

	return s
}

func (state *State) End() {

	bl.AddStabilizeFunc_PreKids(func(){

		leftNode := bl.GetNodeById(state.Z_Left_NodeId)
		rightNode := bl.GetNodeById(state.Z_Right_NodeId)

		oldLeft := rightNode.Left()
		newLeft := leftNode.Left() + leftNode.Width()

		rightNode.SetLeft(newLeft)

		delta := oldLeft - newLeft
		rightNode.SetWidth( rightNode.Width() + delta)
	})
}

