package zindex

import (
	"sort"
	"github.com/amortaza/go-xel"
	"github.com/amortaza/go-bellina"
)

var g_ctxByNodeId map[string] *Ctx

type Ctx struct {
	orderByNodeId map[string] int
	nextOrder int
}

type NodeCtx struct {
	id string
	order int
}

type ByOrder [] *NodeCtx

func (a ByOrder) Len() int {return len(a)}
func (a ByOrder) Swap(i,j int) {a[i],a[j]=a[j],a[i]}
func (a ByOrder) Less(i,j int) bool {return a[i].order < a[j].order}

func (c *Plugin) On(cb func(interface{})) {
	ctx, ok := g_ctxByNodeId[bl.Current_Node.ID]

	if !ok {
		ctx = &Ctx{}
		ctx.orderByNodeId = make(map[string] int)
		g_ctxByNodeId[bl.Current_Node.ID] = ctx

		var order int = 0
		kids := bl.Current_Node.Kids

		for e := kids.Front(); e != nil; e = e.Next() {
		    	kid := e.Value.(*bl.Node)

			ctx.orderByNodeId[kid.ID] = order

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
		node := bl.GetNodeByID(nodectx.id)

		bl.Current_Node.Kids.PushBack(node)

		bl.OnMouseButtonOnNode(node, func(e *bl.MouseButtonEvent) {
			if e.Action == xel.Button_Action_Down {
				ctx.orderByNodeId[node.ID] = ctx.nextOrder
				ctx.nextOrder++
			}
		})
	}
}

