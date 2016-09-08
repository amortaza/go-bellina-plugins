package focus

import (
	"github.com/amortaza/go-bellina"
)

var g_lastNodeId string

var g_onKeyByNodeId map[string] func(interface{})
var g_onLoseFocusByNodeId map[string] func(interface{})
var g_onGainFocusByNodeId map[string] func(interface{})

func newFocusGainLoseEvent(clickedFocusId, loseFocusId string) Event {
	return Event{clickedFocusId, loseFocusId, true, false, nil}
}

func newFocusKeyEvent(hasFocusId string, keyEvent *bl.KeyEvent) Event {
	return Event{hasFocusId, "", false, true, keyEvent}
}


