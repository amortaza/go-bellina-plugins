package focus

import (
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina/event"
	"github.com/amortaza/go-bellina"
)

func (e *Event) Type() string {
	return "focus"
}

func (p *Plugin) Name() string {
	return "focus"
}

func (c *Plugin) GetState() interface{} {
	return nil
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (p *Plugin) Init() {
	g_keyCbByNodeId = make(map[string] func(interface{}))
	g_endCbByNodeId = make(map[string] func(interface{}))

	bl.Plugin( click.NewPlugin() )
	
	event.RegisterLongTerm(bl.EventType_Key, func(e event.Event) {
		if lastNodeID == "" {
			return
		}

		cb, ok := g_keyCbByNodeId[lastNodeID]

		if ok {
			node := bl.GetNodeByID(lastNodeID)
			cb(newEvent(node, e.(*bl.KeyEvent)))
		}
	})

	event.RegisterLongTerm(bl.EventType_Mouse_Button, func(mbe event.Event) {
		if lastNodeID == "" {
			return
		}

		e := mbe.(*bl.MouseButtonEvent)

		if e.Target.ID != lastNodeID {

			endCb, ok := g_endCbByNodeId[lastNodeID]

			if ok {
				endCb(newEvent(e.Target, nil))
			}

			lastNodeID = ""
		}
	})
}

func (p *Plugin) On(cb func(interface{})) {
	p.On2(cb, nil, nil)
}

func (p *Plugin) On2(	cb func(interface{}),
			startCb func(interface{}),
			endCb func(interface{})) {

	bl.On("click", func(i interface{}) {
		e := i.(click.Event)

		if lastNodeID != e.Target.ID {
			if startCb != nil {
				startCb(newEvent(e.Target, nil))
			}
		}

		lastNodeID = e.Target.ID

		g_keyCbByNodeId[lastNodeID] = cb

		if endCb != nil {
			g_endCbByNodeId[lastNodeID] = endCb
		}
	})
}
