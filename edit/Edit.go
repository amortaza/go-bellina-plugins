package edit

import (
	"bellina"
	"plugin/focus"
	"xel"
)

func (c *Plugin) On(cb func(interface{})) {

	editId := "ace"

	shadow := bl.EnsureShadow()
	//editInfo := g_editInfoByEditId[editId]

	bl.Div()
	{
		bl.ID(editId)

		bl.Dim(512, 46)
		bl.Pos(10,10)
		bl.Color(.5,.5,.5)
		bl.Flag(bl.COLOR_SOLID | bl.BORDER_ALL)
		bl.BorderColor(1,1,0)
		bl.BorderThickness(bl.FourOnesInt)
		bl.FontNudge(10,5)
		bl.Font("arial", 11)
		bl.FontColor(1,1,0)
		bl.Label(shadow.Label)

		bl.On("drag", nil)

		bl.On("focus", func(focusEvent interface{}){

			e := focusEvent.(focus.Event)

			if e.KeyEvent.Action == xel.Down {
				shadow.Label += xel.KeyToChar(e.KeyEvent.Key, false, true)
			}
		})
	}
	bl.End()
}

type EditInfo struct {

}

