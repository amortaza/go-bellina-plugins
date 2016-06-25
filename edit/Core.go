package edit

import (
	"github.com/amortaza/go-bellina-plugins/focus"
	"github.com/amortaza/go-xel"
	"math"
	"github.com/amortaza/go-bellina/core"
	"github.com/amortaza/go-bellina"
)

var g_editInfoByEditId map[string] *EditInfo

//type Event struct {
//	X, Y int32
//	Target *bl.Node
//}

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "edit"
}


func (c *Plugin) Init() {
	bl.Plugin( focus.NewPlugin() )

	g_editInfoByEditId = make(map[string] *EditInfo)
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not supported for edit plugin")
}

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}

func insertChar(src string, pos int, char string) string {

	if pos >= len(src) {
		return src + char
	}

	p1 := src[0:pos]
	p2 := src[pos:]

	return p1 + char + p2
}

func backspace(src string, pos int) string {

	if pos < 1 {
		return src
	}

	p1 := src[0:pos-1]
	p2 := src[pos:]

	return p1 + p2
}

func doDelete(src string, pos int) string {

	if pos >= len(src) {
		return src
	}

	p1 := src[0:pos]
	p2 := src[pos+1:]

	return p1 + p2
}

func processKeyDown(key xel.KeyboardKey, alt, ctrl, shift bool, shadow *bl.ShadowNode, editInfo *EditInfo) {
	bt := xel.KeyToBehavior(key, false, true)

	if bt == xel.Key_Behavior_CHAR {
		shadow.Label = insertChar(shadow.Label, editInfo.cursorPos, xel.KeyToChar(key, shift, true))
		editInfo.cursorPos++

	} else {
		if key == xel.Key_HOME {
			editInfo.cursorPos = 0
		}
		if key == xel.Key_END {
			editInfo.cursorPos = len(shadow.Label)
		}
		if key == xel.Key_DELETE {
			shadow.Label = doDelete(shadow.Label, editInfo.cursorPos)
		}
		if key == xel.Key_BACKSPACE {
			shadow.Label = backspace(shadow.Label, editInfo.cursorPos)
			editInfo.cursorPos = int(math.Max(0, float64(editInfo.cursorPos-1)))
		}
		if key == xel.Key_LEFT {
			editInfo.cursorPos = int(math.Max(0, float64(editInfo.cursorPos-1)))
		}
		if key == xel.Key_RIGHT {
			editInfo.cursorPos = int(math.Min(float64(editInfo.cursorPos+1), float64(len(shadow.Label))))
		}
	}
}

func getCursorX(cursorPos int, text string, fontname string, fontsize int32) int32 {
	if cursorPos > len(text) {
		cursorPos = len(text)
	}

	substr := text[:cursorPos]

	g4font := core.GetG4Font(fontname, fontsize)

	return g4font.Width(substr+ " ") + 2
}

func ensureEditInfo(editId string) *EditInfo {
	editInfo, ok := g_editInfoByEditId[editId]

	if !ok {
		editInfo = &EditInfo{}
		g_editInfoByEditId[editId]=editInfo
	}

	return editInfo
}