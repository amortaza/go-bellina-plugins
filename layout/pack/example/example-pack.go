package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-bellina-plugins/layout/pack"
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
			bl.Id("blue")
			bl.Pos(10, 10)
			bl.Dim(500, 500)

			border.Wire(0, 0, 255)

			bl.Div()
			{
				bl.Id("yellow")
				bl.Pos(10, 10)
				bl.Dim(100, 100)

				border.Wire(255, 255, 0)
			}
			bl.End()

			bl.Div()
			{
				bl.Id("green")
				bl.Pos(100, 100)
				bl.Dim(100, 100)

				border.Wire(25, 255, 25)
			}
			bl.End()

			pack.Use()
		}
		bl.End()
	}
	bl.End()
}

func init() {
	runtime.LockOSThread()
}

func main() {

	bl.Start( hal_g5.NewHal(), "Bellina v0.2", 1280, 1024, initialize, tick, nil )

	fmt.Println("bye!")
}


