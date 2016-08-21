package vert

import (
	"github.com/amortaza/go-bellina"
)

// Shared variable across Div()/End()
var g_curState *State

func init() {
	g_stateById = make(map[string] *State)
}

func Id(postfixVertId string) *State {
	vertId := bl.Current_Node.Id + "/" + postfixVertId

	g_curState = ensureState(vertId)

	return g_curState
}

func Use() {

	shadow := bl.EnsureShadow()

	kids := bl.Current_Node.Kids

	for e := kids.Front(); e != nil; e = e.Next() {
	        kid := e.Value.(*bl.Node)

		kidShadow := bl.EnsureShadowById(kid.Id)

		kidShadow.DimWidth__Node_Only()
		kidShadow.DimHeight__Node_Only()

		kidShadow.PosLeft__Node_Only()
		kidShadow.PosTop__Node_Only()
	}

	bl.RegisterShortTerm_LifeCycleTick(func() {
		runLogic(shadow, g_curState)
	})
}
