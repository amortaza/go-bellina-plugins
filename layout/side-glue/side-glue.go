package sideglue

import (
	"github.com/amortaza/go-bellina"
	"fmt"
)

func fake() {
    var _ = fmt.Println
}

// Shared variable across Div()/End()
var g_curState *State

func Id() *State {
	g_curState = ensureState()

	return g_curState
}


func End() {

	state := g_curState

	// to do
	//fmt.Println("+ Pushing Side Glue", )

	bl.AddFunc(func() {
		// to do
		//fmt.Println("--- Processing Side Glue")

		leftNode := bl.GetNodeById(state.Z_Left_NodeId)
		rightNode := bl.GetNodeById(state.Z_Right_NodeId)

		a := rightNode.Left - leftNode.Left
		b := leftNode.Width

		//bl.Disp(leftNode)
		//bl.Disp(rightNode)

		delta := a - b

		//fmt.Println("delta ", delta)

		//fmt.Println("(a) Left Width before ", leftNode.Id, " : ", leftNode.Width)
		leftNode.Width += delta
		//fmt.Println("(b) Left Width after ", leftNode.Id, " : ", leftNode.Width)

		//fmt.Println("--- END Processing Side Glue")
	})

	//fmt.Println("(2) In side glue, left node is ", leftNode.Id, " : " , leftNode.Width, " and parent is ", leftNode.Parent.Id, " : ", leftNode.Parent.Width)
}
