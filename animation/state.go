package animation

import (
	"github.com/amortaza/go-bellina"
	"container/list"
)

var g_states list.List

type AnimState struct {

	tick         func(*bl.ShadowNode, float32)

	NodeId       string
	AnimId       string

	StartValue   float32
	EndValue     float32

	shadow       *bl.ShadowNode

	interpolFunc func() (float32, bool)

	nextPct      func() (float32, bool)
	diff         float32
}


