package horiz_split_pane

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
	"github.com/amortaza/go-bellina-plugins/drag"
)

var g_curState *State

func Use(leftId, handleId, rightId string) {

	var wasNew bool

	g_curState, wasNew = ensureState()

	if wasNew {
		g_curState.leftId, g_curState.handleId, g_curState.rightId = leftId, handleId, rightId
	}

	state := g_curState

	var handleShadow *bl.ShadowNode

	parentId := bl.Current_Node.Id
	parentShadow := bl.EnsureShadowById(parentId)

	sourceLeftWidth := 0

	bl.DivId(leftId)
	{
		sourceLeftWidth = bl.Current_Node.Width()

		docker.Use().AnchorLeft(10).AnchorTop(10).AnchorBottom(10).End()
	}
	bl.End()

	bl.DivId(handleId)
	{
		bl.Left(sourceLeftWidth+10)

		drag_pipe2 := func(x, y int) {
			drag_pipe(x, y, state)
		}

		drag.Use()
		drag.PipeTo(drag_pipe2)

		handleShadow = bl.EnsureShadow()
		handleShadow.Top = 10
		handleShadow.Height = parentShadow.Height - 20
	}
	bl.End()

	bl.DivId(rightId)
	{
		docker_pipe2 := func(x, y, w, h int) {
			docker_pipe(x, y, w, h, state)
		}

		docker.Use().AnchorRight(10).AnchorTop(10).AnchorBottom(10).PipeTo(docker_pipe2).End()

		rightShadow := bl.EnsureShadow()
		oldLeft := rightShadow.Left
		rightShadow.Left = handleShadow.Left + handleShadow.Width
		delta := rightShadow.Left - oldLeft
		rightShadow.Width -= delta
	}
	bl.End()
}

func drag_pipe(x, y int, state *State) {

	leftShadow := bl.EnsureShadowById(state.leftId)
	leftShadow.Width = x - 10

	handleShadow := bl.EnsureShadowById(state.handleId)
	handleShadow.Left = x
}

func docker_pipe(x, y, w, h int, state *State) {

	rightShadow := bl.EnsureShadowById(state.rightId)
	rightShadow.Top = 10

	parentId := rightShadow.BackingNode.Parent.Id
	parentShadow := bl.EnsureShadowById(parentId)
	rightShadow.Height = parentShadow.Height - 20

	handleShadow := bl.EnsureShadowById(state.handleId)
	hright := handleShadow.Left + handleShadow.Width
	rightShadow.Width = parentShadow.Width - hright - 10
}