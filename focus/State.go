package focus

import (
	"github.com/amortaza/go-bellina"
)

var lastNodeID string

var g_keyCbByNodeId map[string] func(interface{})
var g_endCbByNodeId map[string] func(interface{})

type Event struct {
	Target *bl.Node
	KeyEvent *bl.KeyEvent
}

type Plugin struct {
}

