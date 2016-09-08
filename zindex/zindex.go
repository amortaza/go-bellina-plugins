package zindex

import (
	"sort"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-xel2"
)

func On(cb func(interface{})) {
	bl.RequireSettledKids()

	ctx, ok := g_ctxByNodeId[bl.Current_Node.Id]

	if !ok {
		ctx = &Ctx{}
		ctx.orderByNodeId = make(map[string] int)
		g_ctxByNodeId[bl.Current_Node.Id] = ctx

		var order int = 0
		kids := bl.Current_Node.Kids

		for e := kids.Front(); e != nil; e = e.Next() {
		    	kid := e.Value.(*bl.Node)

			ctx.orderByNodeId[kid.Id] = order

			order++
		}

		ctx.nextOrder = order
	}

	var lst [] *NodeCtx

	for nodeid, order := range ctx.orderByNodeId {
		lst = append(lst, &NodeCtx{nodeid, order})
	}

	sort.Sort(ByOrder(lst))

	bl.Current_Node.Kids.Init()

	for _, nodectx := range lst {
		node := bl.GetNodeById(nodectx.id)

		bl.Current_Node.Kids.PushBack(node)

		bl.OnMouseButtonOnNode(node, func(e *bl.MouseButtonEvent) {
			if e.ButtonAction == xel.Button_Action_Down {
				ctx.orderByNodeId[node.Id] = ctx.nextOrder
				ctx.nextOrder++
			}
		})
	}
}



