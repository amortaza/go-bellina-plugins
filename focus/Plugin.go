package focus

import (
	"bellina"
	"plugin/click"
	"fmt"
	"bellina/event"
)

var lastNodeID string

var g_keyCbByNodeId map[string] func(interface{})

type Event struct {
	Target *bl.Node
	KeyEvent *bl.KeyEvent
}

func (e *Event) Type() string {
	return "focus"
}

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "focus"
}

func (c *Plugin) Init() {
	g_keyCbByNodeId = make(map[string] func(interface{}))

	bl.Plugin( click.NewPlugin() )
	
	event.RegisterLongTerm(bl.Key_Event_Type, func(e event.Event) {
		if lastNodeID == "" {
			return
		}

		cb, ok := g_keyCbByNodeId[lastNodeID]

		if ok {
			node := bl.GetNodeByID(lastNodeID)
			cb(newEvent(node, e.(*bl.KeyEvent)))
		}
	})

	event.RegisterLongTerm(bl.Mouse_Button_Event_Type, func(mbe event.Event) {
		if lastNodeID == "" {
			return
		}

		e := mbe.(*bl.MouseButtonEvent)

		if e.Target.ID != lastNodeID {
			lastNodeID = ""
		}
	})
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	fmt.Println("On2 not supported for focusable")
}

func (c *Plugin) On(cb func(interface{})) {

	bl.On("click", func(i interface{}) {
		e := i.(click.Event)

		lastNodeID = e.Target.ID

		g_keyCbByNodeId[lastNodeID] = cb
	})
}

func newEvent(target *bl.Node, keyEvent *bl.KeyEvent) Event {
	return Event{target, keyEvent}
}

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}
