package hover

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-adt"
	"container/list"
)

var g_dirtyCalls = list.New()

// must be interface{}
func On(cb func(interface{})) {

	g_callbacksByNodeId.Add(bl.Current_Node.Id, cb)
}

func AddDirtyCall(cb func()) {
	g_dirtyCalls.PushBack(cb)
}

func init() {
	g_callbacksByNodeId = adt.NewCallbacksById()

	bl.Register_LifeCycle_Before_UserTick_LongTerm(func() {
		g_callbacksByNodeId.ClearAll()
	})

	bl.Register_LifeCycle_After_UserTick_LongTerm(func() {
		for e := g_dirtyCalls.Front(); e != nil; e = e.Next() {

			cb := e.Value.(func())

			cb()
		}

		g_dirtyCalls.Init()
	})

	bl.RegisterLongTerm( bl.EventType_Mouse_Move, func(e bl.Event) {
		event := e.(*bl.MouseMoveEvent)

		curId := event.Target.Id

		if g_lastNodeId != curId {

			inId, outId := curId, g_lastNodeId

			if g_callbacksByNodeId.HasId(inId) {
				eIn := newEvent(inId, outId, true)
				g_callbacksByNodeId.CallAll(inId, eIn)
			}

			if g_callbacksByNodeId.HasId(outId) {
				eOut := newEvent(inId, outId, false)
				g_callbacksByNodeId.CallAll(outId, eOut)
			}

			g_lastNodeId = curId
		}
	})
}

