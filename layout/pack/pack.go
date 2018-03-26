package pack

import (
	"github.com/amortaza/go-bellina"
)

func init() {

	//g_stateByNodeId = make(map[string] *State)
}

//type State struct {
//}

func Use() {

	node := bl.Current_Node

	bl.AddStabilizeFunc_PostKids(func() {

		shadow := bl.EnsureShadowById(node.Id)

		shadow.Width, shadow.Height = runLogic(shadow)
	})
}
