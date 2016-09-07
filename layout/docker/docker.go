package docker

import (
	"github.com/amortaza/go-bellina"
	"fmt"
)

func fake2() {
    var _ = fmt.Println
}

var _ANCHOR_LEFT uint32 = 1 << 0
var _ANCHOR_RIGHT uint32 = 1 << 1
var _ANCHOR_TOP uint32 = 1 << 2
var _ANCHOR_BOTTOM uint32 = 1 << 3

func init() {
	g_stateByNodeId = make(map[string] *State)

	//bl.Register_LifeCycle_AfterUser_Tick(tick)
}

type State struct {
	anchorFlags uint32
}

func Id() (*State) {
	bl.RequireSettledBoundaries()

	return ensureState(bl.Current_Node.Id)
}

func (s *State) AnchorBottom() (*State) {
	s.anchorFlags |= _ANCHOR_BOTTOM

	return s
}

func (s *State) AnchorTop() (*State) {
	s.anchorFlags |= _ANCHOR_TOP

	return s
}

func (s *State) AnchorRight() (*State) {
	s.anchorFlags |= _ANCHOR_RIGHT

	return s
}

func (s *State) AnchorLeft() (*State) {
	s.anchorFlags |= _ANCHOR_LEFT

	return s
}

func (state *State) End() {

	node := bl.Current_Node

	// to do
	//fmt.Println("+ Pushing Docker ", node.Id)

	bl.AddFunc(func() {

		 //to do
		//fmt.Println("* Processing Docker ", node.Id )
		//if node.Id == "<listpane-scroll : listpane>" {
		//	fmt.Println(">> In Docker...", state.anchorFlags & _ANCHOR_RIGHT != 0 && state.anchorFlags & _ANCHOR_LEFT != 0 )
		//}

		left, top, width, height := runLogic(node, state)

		if state.anchorFlags & _ANCHOR_RIGHT != 0 || state.anchorFlags & _ANCHOR_LEFT != 0 {
			node.Left = left
		}

		if state.anchorFlags & _ANCHOR_BOTTOM != 0 || state.anchorFlags & _ANCHOR_TOP != 0 {
			node.Top = top
		}

		if state.anchorFlags & _ANCHOR_RIGHT != 0 && state.anchorFlags & _ANCHOR_LEFT != 0 {
			//to do
			//fmt.Println("In Docker...width ", width)

			node.Width = width
		}

		if state.anchorFlags & _ANCHOR_BOTTOM != 0 && state.anchorFlags & _ANCHOR_TOP != 0 {
			node.Height = height
		}

		// to do
		//bl.Disp(node)
	})
}
