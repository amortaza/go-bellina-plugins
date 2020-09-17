package click

import (
	"github.com/amortaza/go-bellina"
)

type Event struct {
	X, Y int
	Target *bl.Node
}

func On(cb func(interface{})) {
	logic(cb, nil, nil)
}

func On_WithLifeCycle(cb, onDown, onUpAndMiss func(interface{})) {
	logic(cb, onDown, onUpAndMiss)
}

