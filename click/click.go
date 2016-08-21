package click

import (
	"github.com/amortaza/go-bellina"
)

type Event struct {
	X, Y int
	Target *bl.Node
}

// click.On( func )
func On(cb func(interface{})) {
	logic(cb, nil, nil)
}

// click.On_WithLifeCycle( func )
func On_WithLifeCycle(cb func(interface{}), onDown func(interface{}), onUpAndMiss func(interface{})) {
	logic(cb, onDown, onUpAndMiss)
}

