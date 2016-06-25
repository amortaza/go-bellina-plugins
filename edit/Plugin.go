package edit

import (
	"bellina"
	"plugin/focus"
	"xel"
	"math"
)

var a float64

func (c *Plugin) Tick() {
	editInfo, ok := g_editInfoByEditId["ace"]

	if ok {
		if !editInfo.hasFocus {
			editInfo.opacity = 0
			return
		}

		editInfo.opacity = (float32(math.Sin(a)) + 1 ) / 2 + .5
	}

	a += .3
}

type EditInfo struct {
	cursorPos int
	opacity float32
	hasFocus bool
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) On(cb func(interface{})) {

	editId := "ace"

	shadow := bl.EnsureShadow()
	editInfo := ensureEditInfo(editId)

	var fontheight int32

	bl.Div()
	{
		bl.ID(editId)

		bl.Pos(10,10)
		bl.Color(.5,.5,.5)
		bl.Flag(bl.Z_COLOR_SOLID | bl.Z_BORDER_ALL)
		bl.BorderColor(1,1,0)
		bl.BorderThickness(bl.FourOnesInt)
		bl.FontNudge(10,5)
		bl.Font("tahoma", 7)
		bl.FontColor(1,1,0)
		bl.Label(shadow.Label)

		fontheight =  bl.GetFontHeight() + 4
		bl.Dim(512, fontheight )

		bl.On2("focus", func(focusEvent interface{}){

			e := focusEvent.(focus.Event)


			if e.KeyEvent.Action == xel.Button_Action_Down {
				key := e.KeyEvent.Key
				processKeyDown(key, e.KeyEvent.Alt, e.KeyEvent.Ctrl, e.KeyEvent.Shift, shadow, editInfo)
			}

		}, func(focusEvent interface{}) {
			editInfo.hasFocus = true
			editInfo.cursorPos = len(shadow.Label)

		}, func(focusEvent interface{}) {
			editInfo.hasFocus = false

		})

		parent := bl.Current_Node

		bl.Div()
		{
			bl.ID(editId+":cursoer")
			bl.Dim(2, fontheight)
			bl.Pos(getCursorX(editInfo.cursorPos, shadow.Label, parent.FontName, parent.FontSize),0)
			bl.Color(1,.95,.95)
			bl.NodeOpacity1f(editInfo.opacity)
		}
		bl.End()
	}
	bl.End()
}


