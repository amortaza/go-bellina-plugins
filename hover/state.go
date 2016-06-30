package hover

import "github.com/amortaza/go-adt"

var callbacksByNodeId *adt.CallbacksByID
var lastNodeId string

type Event struct {
	InNodeId  string
	OutNodeId string
	IsInEvent bool
}

func newEvent(inNodeId, outNodeId string, isInEvent bool) *Event {
	c := &Event{inNodeId, outNodeId, isInEvent}

	return c
}


