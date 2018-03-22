package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-bellina-plugins/vert-split-pane"
)

func initialize() {
	go_dark_ux.Init()
}

func tick() {

	bl.Root()
	{
		bl.Pos(0,0)

			bl.Div()
			{
				bl.Id("yellow")
				bl.Pos(10, 10)
				bl.Dim(1004, 748)

				bl.Div()
				{
					bl.Id("top")
					bl.Height(100)
					border.Wire(0, 0, 255)
				}
				bl.End()

				bl.Div()
				{
					bl.Id("handle")
					bl.Height(50)
					border.Wire(0, 0, 0)
				}
				bl.End()

				bl.Div()
				{
					bl.Id("bottom")

					border.Wire(0, 255, 0)
				}
				bl.End()

				border.Wire(255, 255, 0)

				vert_split_pane.Use("top", "handle", "bottom")
			}
			bl.End()

		border.Wire(255, 0, 0)
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


