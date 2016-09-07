package sideglue

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

func (s *State) End() {
	End()
}

