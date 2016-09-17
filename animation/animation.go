package animation

import (
	"github.com/amortaza/go-bellina"
	"container/list"
)

func init() {
	bl.Register_LifeCycle_After_UserTick_LongTerm(tick)
}

func StartPath(nodeId string, animId string, startValue, endValue float32, numSteps int, cb func(shadow *bl.ShadowNode, value float32)) {
	state := &AnimState{NodeId: nodeId, AnimId: animId, shadow: bl.EnsureShadowById(nodeId)}

	g_states.PushBack(state)

	state.tick = cb
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

	state.interpolFunc = Linear(state)
}

func tick() {
	var elements list.List

	for e := g_states.Front(); e != nil; e = e.Next() {
	        state := e.Value.(*AnimState)

		value, valid := state.interpolFunc()

		if valid {
			state.tick(state.shadow, value)

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

