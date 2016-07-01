package animation

import (
)
import (
	"github.com/amortaza/go-bellina"
	"container/list"
)

var plugin *Plugin

func StartPath(nodeId string, animId string, startValue, endValue float32, numSteps int, interpolMethod string, cb func(shadow *bl.ShadowNode, value float32)) {
	state := &AnimState{NodeId: nodeId, AnimId: animId}
	g_states.PushBack(state)

	state.Tick  = cb
	state.StartValue = startValue
	state.EndValue = endValue

	state.diff = endValue - startValue

	var curPct float32 = 0
	var deltaPct float32 = 1.0 / float32(numSteps)
	remaining := numSteps

	state.nextPct = func() (float32, bool) {

		if remaining <= 0 {
			return -1, false
		}

		ret := curPct
		curPct += deltaPct
		remaining--

		if remaining == 0 {
			ret = 1
		}

		return ret, true
	}

	state.InterpolFunc = Linear(state)
}

func (c *Plugin) Tick() {
	elements := list.New()

	for e := g_states.Front(); e != nil; e = e.Next() {
	        state := e.Value.(*AnimState)

		shadow := bl.EnsureShadowById(state.NodeId)
		value, valid := state.InterpolFunc()

		if valid {
			state.Tick(shadow, value)
		} else {
			elements.PushBack(e)
		}
	}

	if elements.Len() > 0 {
		for e := elements.Front(); e != nil; e = e.Next() {
		        el := e.Value.(*list.Element)
			g_states.Remove(el)
		}
	}
}

func Linear(state *AnimState) (func() (float32, bool)) {
	return func() (float32, bool) {
		pct, valid := state.nextPct()

		return state.StartValue + state.diff * pct, valid
	}
}