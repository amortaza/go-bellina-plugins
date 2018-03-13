package horiz_split_pane

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
	"github.com/amortaza/go-bellina-plugins/drag"
)

func Use(leftId, handleId, rightId string) {

	bl.DivId(leftId)
	{
		docker.Use().AnchorLeft(10).AnchorTop(10).AnchorBottom(10).End()
	}
	bl.End()

	bl.DivId(handleId)
	{
		drag.Use()
	}
	bl.End()

	bl.DivId(rightId)
	{
		docker.Use().AnchorRight(10).AnchorTop(10).AnchorBottom(10).End()
	}
	bl.End()
}
