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

	bl.Register_LifeCycle_AfterUser_Tick(tick)
}

type State struct {
	anchorFlags uint32
}

func Id() (*State) {
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

func (s *State) End() {
	shadow := bl.EnsureShadow()

	shadow.PosLeft__Node_Only()
	shadow.PosTop__Node_Only()
	shadow.DimWidth__Node_Only()
	shadow.DimHeight__Node_Only()
}

func tick() {
	for key, state := range g_stateByNodeId {
		shadow := bl.EnsureShadowById(key)

		runLogic(shadow, state)
	}
}
