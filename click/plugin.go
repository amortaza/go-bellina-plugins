package click

import (
	"github.com/amortaza/go-bellina/event"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-xel"
	"fmt"
)

var NAME = "click"

var plugin *Plugin

type Plugin struct {
}

func (c *Plugin) Name() string {
	return NAME
}

func (c *Plugin) GetState() interface{} {
	return nil
}

func (c *Plugin) Init() {
}

func (c *Plugin) Tick() {
}

func (c *Plugin) Reset() {
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On(cb func(interface{})) {
	c.On2(cb, nil, nil)
}

func (c *Plugin) On2(cb func(interface{}), onDown func(interface{}), onUpAndMiss func(interface{})) {

	event.RegisterShortTerm(bl.EventType_Mouse_Button, func(event event.Event) {

		e := event.(*bl.MouseButtonEvent)

		if e.ButtonAction == xel.Button_Action_Down {

			lastNodeID = e.Target.Id

			if onDown != nil {
				onDown(Event{bl.Mouse_X, bl.Mouse_X, e.Target})
			}

		} else if e.ButtonAction == xel.Button_Action_Up {

			if lastNodeID == e.Target.Id {
				// we have a click!
				cb(Event{bl.Mouse_X, bl.Mouse_X, e.Target})

			} else {
				lastNodeID = ""

				if onUpAndMiss != nil {
					onUpAndMiss(nil)
				}
			}

		} else {
			fmt.Println("Button action not recognized in click.Plugin")
		}
	})
}

func NewPlugin() *Plugin {
	plugin = &Plugin{}

	return plugin
}
