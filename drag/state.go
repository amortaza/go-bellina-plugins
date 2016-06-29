package drag

import "github.com/amortaza/go-bellina"

func newEvent(target *bl.Node) Event {
	return Event{
		bl.Mouse_X, bl.Mouse_X,
		target,
	}
}

