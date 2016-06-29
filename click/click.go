package click

import (
	"github.com/amortaza/go-bellina"
)

type Event struct {
	X, Y int32
	Target *bl.Node
}

func On(cb func(interface{})) {
	plugin.On(cb)
}

func On2(cb func(interface{}), startCb func(interface{}), endCb func(interface{})) {
	plugin.On2(cb, startCb, endCb)
}

