package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-bellina-plugins/focus"
)

func initialize() {
	go_dark_ux.Init()
}

func tick() {

	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(800, 700)

		bl.Div()
		{
			bl.Id("red")
			bl.Pos(10,10)
			bl.Dim(100,100)

			border.Wire()

			focus.On_LifeCycle(func(i interface{}) {
				fmt.Println("Key")
			}, func(i interface{}) {
				fmt.Println("Gain")
			}, func(i interface{}) {
				fmt.Println("Lose")
			})
		}
		bl.End()
		bl.Div()
		{
			bl.Id("green")
			bl.Pos(100,100)
			bl.Dim(100,100)

			border.Wire()

			focus.On_LifeCycle(func(i interface{}) {
				fmt.Println("Key")
			}, func(i interface{}) {
				fmt.Println("Gain")
			}, func(i interface{}) {
				fmt.Println("Lose")
			})
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
	bl.Start( haloob.New(), 1280, 1024, "Bellina v0.2", initialize, tick, uninit )

	fmt.Println("bye!")
}


