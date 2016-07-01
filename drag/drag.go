package drag

import "github.com/amortaza/go-bellina"

type Event struct {
	X, Y int
	Target *bl.Node
}

func On(cb func(interface{})) {
	plugin.On(cb)
}

