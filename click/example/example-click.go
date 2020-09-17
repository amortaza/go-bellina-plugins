package main

import (
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-hal-oob"
	"runtime"
)

func initialize() {
	go_dark_ux.Init()
}

func tick() {
	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(1024, 768)
		border.Fill(50, 0, 0)

		bl.Div()
		{
			bl.Id("green")
			bl.Pos(10,10)
			bl.Dim(100,100)

			border.Fill(0, 50, 0)

			click.On(func(e interface{}) {
				clickEvent := e.(click.Event)

				fmt.Println("Clicked on", clickEvent.Target.Id)
				fmt.Println("Clicked at", clickEvent.X, clickEvent.Y)
			})
		}
		bl.End()

		bl.Div()
		{
			bl.Id("blue")
			bl.Pos(110, 110)
			bl.Dim(100, 100)

			border.Fill(0, 0, 70)

			click.On_WithLifeCycle(

				// click success!
				func(e interface{}) {
					clickEvent := e.(click.Event)

					fmt.Println("Clicked on", clickEvent.Target.Id, clickEvent.X, clickEvent.Y)
				},

				// click initiated
				func(e interface{}) {
					clickEvent := e.(click.Event)

					fmt.Println("Mouse Down", clickEvent.Target.Id, clickEvent.X, clickEvent.Y)
				},

				// click failed
				func(e interface{}) {
					// e will be nil

					fmt.Println("Click miss")
				},
			)
		}
		bl.End()
	}
	bl.End()
}

func uninit() {
}

func init() {
	runtime.LockOSThread()
}

func main() {
	bl.Start( haloob.NewHal(), "i3wmfloater",1200, 100,1280, 1024, initialize, tick, uninit )

	fmt.Println("bye!")
}


