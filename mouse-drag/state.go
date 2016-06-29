package mouse_drag

import "github.com/amortaza/go-bellina"

var lastNodeID string
var startX, startY int32
var mouseOffsetX, mouseOffsetY int32

func newEvent(mouseX, mouseY int32, target *bl.Node) Event {
	return Event{
		mouseX, mouseY,
		target,
		startX, startY,
		mouseOffsetX, mouseOffsetY,
	}
}

