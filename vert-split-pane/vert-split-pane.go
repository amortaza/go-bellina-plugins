package vert_split_pane

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
	"github.com/amortaza/go-bellina-plugins/drag"
)

var g_curState *State

func Use(topId, handleId, bottomId string) {

	var wasNew bool

	g_curState, wasNew = ensureState()

	if wasNew {
		g_curState.topId, g_curState.handleId, g_curState.bottomId = topId, handleId, bottomId
	}

	state := g_curState

	var handleShadow *bl.ShadowNode

	parentId := bl.Current_Node.Id
	parentShadow := bl.EnsureShadowById(parentId)

	bl.DivId(topId)
	{
		docker.Use().AnchorLeft(10).AnchorTop(10).AnchorRight(10).End()
	}
	bl.End()

	bl.DivId(handleId)
	{
		bl.Left(20)
		bl.Top(20)

		drag_pipe2 := func(x, y int) {
			drag_pipe(x, y, state)
		}

		drag.Use()
		drag.PipeTo(drag_pipe2)

		handleShadow = bl.EnsureShadow()
		handleShadow.Left = 10
		handleShadow.Width = parentShadow.Width - 20
	}
	bl.End()

	bl.DivId(bottomId)
	{
		docker_pipe2 := func(x, y, w, h int) {
			docker_pipe(x, y, w, h, state)
		}

		docker.Use().AnchorLeft(10).AnchorBottom(10).AnchorRight(10).PipeTo(docker_pipe2).End()

		bottomShadow := bl.EnsureShadow()
		oldBottom := bottomShadow.Left
		bottomShadow.Top = handleShadow.Top + handleShadow.Height
		delta := bottomShadow.Top - oldBottom
		bottomShadow.Height -= delta
	}
	bl.End()
}

func drag_pipe(x, y int, state *State) {

	topShadow := bl.EnsureShadowById(state.topId)
	topShadow.Height = y - 10

	handle := bl.EnsureShadowById(state.handleId)
	handle.Top = y
}

func docker_pipe(x, y, w, h int, state *State) {

	bottom := bl.EnsureShadowById(state.bottomId)
	bottom.Left = 10

	parentId := bottom.BackingNode.Parent.Id
	parent := bl.EnsureShadowById(parentId)
	bottom.Width = parent.Width - 20

	handle := bl.EnsureShadowById(state.handleId)
	hbottom := handle.Top + handle.Height
	bottom.Height = parent.Height - hbottom - 10
}