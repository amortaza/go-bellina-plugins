package click

import (
	"github.com/amortaza/go-bellina"
)

type Event struct {
	X, Y int
	Target *bl.Node
}

// click.On( cb )
func On(cb func(interface{})) {
	logic(cb, nil, nil)
}

// click.On_WithLifeCycle( successCb, downCb, upAndMissCb )
func On_WithLifeCycle(cb, onDown, onUpAndMiss func(interface{})) {
	logic(cb, onDown, onUpAndMiss)
}

