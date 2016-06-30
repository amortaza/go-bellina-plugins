package focus

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina/event"
	"github.com/amortaza/go-xel"
)

var plugin *Plugin

type Plugin struct {
}

func NewPlugin() *Plugin {
	plugin = &Plugin{}

	return plugin
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) Tick() {
}

func (c *Plugin) Reset() {
}

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

func (p *Plugin) On(cb func(interface{})) {
	p.On2(cb, nil, nil)
}

func (p *Plugin) Init() {
	g_onKeyByNodeId = make(map[string] func(interface{}))
	g_onLoseFocusByNodeId = make(map[string] func(interface{}))
	g_onGainFocusByNodeId = make(map[string] func(interface{}))

	bl.Plugin( click.NewPlugin() )

	event.RegisterLongTerm(bl.EventType_Key, func(e event.Event) {
		if lastNodeId == "" {
			return
		}

		onKey, ok := g_onKeyByNodeId[lastNodeId]

		if ok {
			onKey(newFocusKeyEvent(lastNodeId, e.(*bl.KeyEvent)))
		}
	})

	event.RegisterLongTerm(bl.EventType_Mouse_Button, func(mbe event.Event) {
		if lastNodeId == "" {
			return
		}

		e := mbe.(*bl.MouseButtonEvent)

		if e.ButtonAction == xel.Button_Action_Down {
			return
		}

		if e.Target.Id != lastNodeId {

			onLoseFocus, ok := g_onLoseFocusByNodeId[lastNodeId]

			if ok {
				onLoseFocus(newFocusGainLoseEvent(e.Target.Id, lastNodeId))
			}

			newId := e.Target.Id

			onGainFocus, ok2 := g_onGainFocusByNodeId[newId]

			if ok2 {
				onGainFocus(newFocusGainLoseEvent(newId, lastNodeId))
			}

			lastNodeId = newId
		}
	})
}

func (p *Plugin) On2(	onKey func(interface{}),
			onGainFocus func(interface{}),
			onLoseFocus func(interface{})) {

	if plugin == nil {
		panic("You did not load the focus plugin")
	}

	nodeId := bl.Current_Node.Id

	click.On( func(i interface{}) {
		e := i.(click.Event)

		if lastNodeId != e.Target.Id {
			if onGainFocus != nil && e.Target.Id == nodeId {
				onGainFocus(newFocusGainLoseEvent(nodeId, lastNodeId))
			}
		}

		if nodeId == e.Target.Id {
			lastNodeId = e.Target.Id
		}
	})

	if onKey != nil {
		g_onKeyByNodeId[nodeId] = onKey
	}

	if onLoseFocus != nil {
		g_onLoseFocusByNodeId[nodeId] = onLoseFocus
	}

	if onGainFocus != nil {
		g_onGainFocusByNodeId[nodeId] = onGainFocus
	}
}
