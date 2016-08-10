package vert

import (
//	"github.com/amortaza/go-bellina"
)

func init() {
	plugin = &Plugin{}
//	bl.Plugin(plugin)
}
/*
func SetSpacing(spacing int) {
	bl.SetI( "vert", "spacing", spacing )
}

func SetPercent(percent int) {
	bl.SetI( "vert", "percent", percent )
}

func FillRemaining() {
	bl.SetI( "vert", "percent", -1 )
}
*/
func Use() {
	plugin.On(nil)
}
