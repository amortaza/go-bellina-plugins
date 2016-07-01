package focus

import (
	"github.com/amortaza/go-bellina"
)

type Event struct {
	ClickedNodeId string
	LoseFocusNodeId string
	IsGainFocusEvent bool
	IsKeyEvent bool
	KeyEvent *bl.KeyEvent
}

func On(cb func(interface{})) {
	plugin.On(cb)
}

func On2(cb func(interface{}), startCb func(interface{}), endCb func(interface{})) {
	plugin.On2(cb, startCb, endCb)
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}
