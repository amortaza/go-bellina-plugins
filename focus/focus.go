package focus

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-xel2"
)

type Event struct {
	FocusNodeId      string
	LoseFocusNodeId  string
	IsGainFocusEvent bool
	IsKeyEvent       bool
	KeyEvent         *bl.KeyEvent
}

func init() {
	g_onKeyByNodeId = make(map[string] func(interface{}))
	g_onLoseFocusByNodeId = make(map[string] func(interface{}))
	g_onGainFocusByNodeId = make(map[string] func(interface{}))

	bl.Register_LifeCycle_Init(onBlInit)
}

func On(onKey func(interface{})) {
	On_LifeCycle(onKey, nil, nil)
}

func On_LifeCycle(onKey func(interface{}), onGainFocus func(interface{}), onLoseFocus func(interface{})) {

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

func onBlInit() {
	bl.RegisterLongTerm(bl.EventType_Key, func(e bl.Event) {
		if lastNodeId == "" {
			return
		}

		onKey, ok := g_onKeyByNodeId[lastNodeId]

		if ok {
			onKey(newFocusKeyEvent(lastNodeId, e.(*bl.KeyEvent)))
		}
	})

	bl.RegisterLongTerm(bl.EventType_Mouse_Button, func(mbe bl.Event) {
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