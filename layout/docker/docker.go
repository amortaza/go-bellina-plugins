package docker

import (
	"github.com/amortaza/go-bellina"
)

var Z_ANCHOR_LEFT uint32 = 1 << 0
var Z_ANCHOR_RIGHT uint32 = 1 << 1
var Z_ANCHOR_TOP uint32 = 1 << 2
var Z_ANCHOR_BOTTOM uint32 = 1 << 3

type State struct {
	anchorFlags uint32
}

func AnchorBottom() {
	s := ensureState(bl.Current_Node.Id)
	s.anchorFlags |= Z_ANCHOR_BOTTOM
}

func AnchorTop() {
	s := ensureState(bl.Current_Node.Id)
	s.anchorFlags |= Z_ANCHOR_TOP
}

func AnchorRight() {
	s := ensureState(bl.Current_Node.Id)
	s.anchorFlags |= Z_ANCHOR_RIGHT
}

func AnchorLeft() {
	s := ensureState(bl.Current_Node.Id)
	s.anchorFlags |= Z_ANCHOR_LEFT
}

func Use() {
	plugin.On(nil)
}