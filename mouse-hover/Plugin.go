package mouse_hover

import (
	"github.com/amortaza/go-bellina/event"
	"github.com/amortaza/go-adt"
	"github.com/amortaza/go-bellina"
)

var g_callbacksByNodeID *adt.CallbacksByID
var lastNodeID string

type Event struct {
	InNodeID string
	OutNodeID string
	IsInEvent bool
}

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "hover"
}

func (c *Plugin) GetState() interface{} {
	return nil
}

func (c *Plugin) Init() {
	g_callbacksByNodeID = adt.NewCallbacksByID()

	event.RegisterLongTerm( bl.EventType_Mouse_Move, func(e event.Event) {
		event := e.(*bl.MouseMoveEvent)

		currentNode := event.Target

		if lastNodeID != currentNode.Id {

			inID, outID := currentNode.Id, lastNodeID

			if g_callbacksByNodeID.HasId(inID) {
				eIn := NewEvent(inID, outID, true)
				g_callbacksByNodeID.CallAll(inID, eIn)
			}

			if g_callbacksByNodeID.HasId(outID) {
				eOut := NewEvent(outID, inID, false)
				g_callbacksByNodeID.CallAll(outID, eOut)
			}

			lastNodeID = currentNode.Id
		}
	})
}

func (c *Plugin) Tick() {
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On(cb func(interface{})) {
	g_callbacksByNodeID.Add(bl.Current_Node.Id, cb)
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not supoorted in mouse_hover.Plugin")
}

func NewEvent(inNodeID, outNodeID string, isInEvent bool) *Event {
	c := &Event{inNodeID, outNodeID, isInEvent}

	return c
}

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}
