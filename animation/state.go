package animation

import (
	"github.com/amortaza/go-bellina"
	"container/list"
)

type AnimState struct {

	Tick func(*bl.ShadowNode, float32)

	NodeId string
	AnimId string

	StartValue float32
	EndValue float32

	InterpolFunc func() (float32, bool)

	nextPct func() (float32, bool)
	diff float32
}

var g_states *list.List
