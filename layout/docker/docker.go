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
	topPadding, leftPadding, rightPadding, bottomPadding int
	pipeTo func(left, top, width, height int)
}

func Use() (*State) {

	return ensureState(bl.Current_Node.Id)
}

func (state *State) PipeTo( pipeTo func(left, top, width, height int) ) (*State) {

	state.pipeTo = pipeTo

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

		shadow := bl.EnsureShadowById(node.Id)

		left, top, width, height := runLogic(shadow, state)

		if state.pipeTo == nil {

			if state.anchorFlags&_ANCHOR_RIGHT != 0 || state.anchorFlags&_ANCHOR_LEFT != 0 {
				shadow.Left = left
			}

			if state.anchorFlags&_ANCHOR_BOTTOM != 0 || state.anchorFlags&_ANCHOR_TOP != 0 {
				shadow.Top = top
			}

			if state.anchorFlags&_ANCHOR_RIGHT != 0 && state.anchorFlags&_ANCHOR_LEFT != 0 {
				shadow.Width = width
			}

			if state.anchorFlags&_ANCHOR_BOTTOM != 0 && state.anchorFlags&_ANCHOR_TOP != 0 {
				shadow.Height = height
			}

		} else {

			state.pipeTo(left, top, width, height)
		}
	})
}
