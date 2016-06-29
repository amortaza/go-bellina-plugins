package hover

import "github.com/amortaza/go-adt"

var g_callbacksByNodeID *adt.CallbacksByID
var lastNodeID string

type Event struct {
	InNodeID string
	OutNodeID string
	IsInEvent bool
}

func newEvent(inNodeID, outNodeID string, isInEvent bool) *Event {
	c := &Event{inNodeID, outNodeID, isInEvent}

	return c
}


