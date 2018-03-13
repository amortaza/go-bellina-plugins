package docker

import (
	"github.com/amortaza/go-bellina"
)

var _ANCHOR_LEFT uint32 = 1 << 0
var _ANCHOR_RIGHT uint32 = 1 << 1
var _ANCHOR_TOP uint32 = 1 << 2
var _ANCHOR_BOTTOM uint32 = 1 << 3

func init() {

	g_stateByNodeId = make(map[string] *State)
}

type State struct {
	anchorFlags uint32
	sudo string
	topPadding, leftPadding, rightPadding, bottomPadding int
}

func Use() (*State) {

	bl.RequireSettledBoundary()

	return ensureState(bl.Current_Node.Id)
}

func (state *State) Sudo(sudo string) (*State) {

	state.sudo = sudo

	return state
}

func (state *State) AnchorBottom(padding int) (*State) {

	state.anchorFlags |= _ANCHOR_BOTTOM

	state.bottomPadding = padding

	return state
}

func (state *State) AnchorTop(padding int) (*State) {

	state.anchorFlags |= _ANCHOR_TOP

	state.topPadding = padding

	return state
}

func (state *State) AnchorRight(padding int) (*State) {

	state.anchorFlags |= _ANCHOR_RIGHT

	state.rightPadding = padding

	return state
}

func (state *State) AnchorLeft(padding int) (*State) {

	state.anchorFlags |= _ANCHOR_LEFT

	state.leftPadding = padding

	return state
}

func (state *State) End() {

	node := bl.Current_Node

	bl.AddStabilizeFunc_PreKids(func() {

		left, top, width, height := runLogic(node, state)

		if state.anchorFlags & _ANCHOR_RIGHT != 0 || state.anchorFlags & _ANCHOR_LEFT != 0 {
			node.SetOwnerOfLeft( state.sudo )
			node.SetLeft(left)
		}

		if state.anchorFlags & _ANCHOR_BOTTOM != 0 || state.anchorFlags & _ANCHOR_TOP != 0 {
			node.SetOwnerOfTop(state.sudo)
			node.SetTop(top)
		}

		if state.anchorFlags & _ANCHOR_RIGHT != 0 && state.anchorFlags & _ANCHOR_LEFT != 0 {
			node.SetOwnerOfWidth(state.sudo)
			node.SetWidth(width)
		}

		if state.anchorFlags & _ANCHOR_BOTTOM != 0 && state.anchorFlags & _ANCHOR_TOP != 0 {
			node.SetOwnerOfHeight(state.sudo)
			node.SetHeight(height)
		}
	})
}
