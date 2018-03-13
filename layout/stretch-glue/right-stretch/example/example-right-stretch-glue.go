package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-bellina-plugins/drag"
	"github.com/amortaza/go-bellina-plugins/layout/stretch-glue/right-stretch"
)

func initialize() {
	go_dark_ux.Init()
}

func tick() {

	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(1024, 768)

		bl.Div()
		{
			bl.Id("red")
			bl.Pos(10, 10)
			bl.Dim(200, 200)
			bl.SettleBoundary()

			border.Wire(255, 0, 0)
			drag.Use()
		}
		bl.End()

		bl.Div()
		{
			bl.Id("green")
			bl.Pos(300, 300)
			bl.Dim(300, 300)

			drag.Use()
			//stretch_stick.StickTo("red").AnchorLeft().End()

			border.Wire(0, 255, 0)
		}
		bl.End()

		stretch_glue_right.Id().LeftNodeId("red").RightNodeId("green").End()
	}
	bl.End()
}

func init() {
	runtime.LockOSThread()
}

func main() {
	bl.Start( hal_g5.NewHal(), "Bellina v0.2", 1024, 768, initialize, tick, nil )

	fmt.Println("bye!")
}


