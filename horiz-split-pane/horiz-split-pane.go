package horiz_split_pane

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
	"github.com/amortaza/go-bellina-plugins/drag"
)

func Use(leftId, handleId, rightId string) {

	var handleShadow *bl.ShadowNode

	parentId := bl.Current_Node.Id
	parent := bl.EnsureShadowById(parentId)

	bl.DivId(leftId)
	{
		//bl.SettleBoundary()

		docker.Use().AnchorLeft(10).AnchorTop(10).AnchorBottom(10).End()
	}
	bl.End()

	bl.DivId(handleId)
	{
		bl.Left(110)
		bl.Top(10)

		drag.Use()
		drag.PipeTo(drag_pipe)

		handleShadow = bl.EnsureShadow()
		handleShadow.Top = 10
		handleShadow.Height = parent.Height - 20
	}
	bl.End()

	bl.DivId(rightId)
	{
		//bl.SettleBoundary()

		docker.Use().AnchorRight(10).AnchorTop(10).AnchorBottom(10).PipeTo(docker_pipe).End()

		rightShadow := bl.EnsureShadow()
		oldLeft := rightShadow.Left
		rightShadow.Left = handleShadow.Left + handleShadow.Width
		delta := rightShadow.Left - oldLeft
		rightShadow.Width -= delta
	}
	bl.End()
}

func drag_pipe(x, y int) {

	left := bl.EnsureShadowById("left")
	left.Width = x - 10

	handle := bl.EnsureShadowById("handle")
	handle.Left = x
}

func docker_pipe(x, y, w, h int) {

	right := bl.EnsureShadowById("right")
	right.Top = 10

	parentId := right.BackingNode.Parent.Id
	parent := bl.EnsureShadowById(parentId)
	right.Height = parent.Height - 20


	handle := bl.EnsureShadowById("handle")
	hright := handle.Left + handle.Width
	right.Width = parent.Width - hright - 10
}