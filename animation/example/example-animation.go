package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina-plugins/animation"
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
			bl.Pos(10,10)
			bl.Dim(100,100)

			border.Wire(255, 255, 0)

			click.On(func(e interface{}) {
				animation.StartPath("red", "hi", 0, 400, 200, func(shadow *bl.ShadowNode, value float32) {
					shadow.SetLeft_on_Self_and_Node(int(value), "animation")
				})
			})

			shadow := bl.EnsureShadow()
			shadow.SetLeft_on_Node_Only("animation")
		}
		bl.End()

		bl.Div()
		{
			bl.Id("green")
			bl.Pos(100,100)
			bl.Dim(100,100)

			border.Wire(255, 255, 0)
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
	bl.Start( hal_g5.NewHal(), "Bellina v0.2", 1280, 1024, initialize, tick, uninit )

	fmt.Println("bye!")
}


