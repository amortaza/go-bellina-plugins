package mouse_drag

import "github.com/amortaza/go-bellina"

var gLastNodeId string
var gStartX, gStartY int
var gMouseOffsetX, gMouseOffsetY int
var gEndCbByNodeId map[string] func(interface{})

func newEvent(mouseX, mouseY int, target *bl.Node) Event {
	return Event{
		mouseX, mouseY,
		target,
		gStartX, gStartY,
		gMouseOffsetX, gMouseOffsetY,
	}
}

