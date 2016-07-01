package mouse_drag

import "github.com/amortaza/go-bellina"

var gLastNodeId string
var startX, startY int
var mouseOffsetX, mouseOffsetY int

func newEvent(mouseX, mouseY int, target *bl.Node) Event {
	return Event{
		mouseX, mouseY,
		target,
		startX, startY,
		mouseOffsetX, mouseOffsetY,
	}
}

