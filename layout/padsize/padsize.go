package padsize

import (
	"github.com/amortaza/go-bellina"
)

func Id() (*State) {
	bl.RequireSettledBoundary()
	bl.RequireSettledKids()

	state := ensureState(bl.Current_Node.Id)

	return state
}

func runLogic(node *bl.Node, state *State) {

	bl.AddFunc( func() {
		padLeft := state.Z_Left
		padTop := state.Z_Top
		padRight := state.Z_Right
		padBottom := state.Z_Bottom

		for e := node.Kids.Front(); e != nil; e = e.Next() {
			kid := e.Value.(*bl.Node)

			if kid.left < padLeft {
				//kid.OwnLeft("pad")
				kid.left = padLeft
			}

			if kid.top < padTop {
				//kid.OwnTop("pad")
				kid.top = padTop
			}

			right := kid.left + kid.width - 1

			if right > node.width- padRight {
				//kid.OwnWidth("pad")
				kid.width = node.width - padRight - kid.left
			}

			bottom := kid.top + kid.height - 1

			if bottom > node.height- padBottom {
				//kid.OwnHeight("pad")
				kid.height = node.height - padBottom - kid.top
			}
		}
	})
}

