package vert

import (
	"github.com/amortaza/go-bellina"
)

// Shared variable across Div()/End()
var g_curState *State

func init() {
	g_stateById = make(map[string] *State)
}

func Id() *State {
	g_curState = ensureState(bl.Current_Node.Id)

	return g_curState
}

func Use() {

	shadow := bl.EnsureShadow()

	kids := bl.Current_Node.Kids

	for e := kids.Front(); e != nil; e = e.Next() {
	        kid := e.Value.(*bl.Node)

		kidShadow := bl.EnsureShadowById(kid.Id)

		kidShadow.PosTop__Node_Only()
	}

	bl.Register_LifeCycle_AfterUser_Tick_ShortTerm(func() {
		runLogic(shadow, g_curState)
	})
}
