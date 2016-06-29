package vert

import (
	"github.com/amortaza/go-bellina"
)

type State struct {
}

func SetSpacing(spacing int32) {
	bl.SetI( "vert", "spacing", spacing )
}

func SetPercent(percent int32) {
	bl.SetI( "vert", "percent", percent )
}

func FillRemaining() {
	bl.SetI( "vert", "percent", -1 )
}

func Use() {
	plugin.On(nil)
}
