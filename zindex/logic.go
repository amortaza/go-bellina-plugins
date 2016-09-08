package zindex

var g_ctxByNodeId map[string] *Ctx

type Ctx struct {
	orderByNodeId map[string] int
	nextOrder int
}

type NodeCtx struct {
	id string
	order int
}

func init() {
	g_ctxByNodeId = make(map[string] *Ctx)
}

type ByOrder [] *NodeCtx

func (a ByOrder) Len() int {return len(a)}
func (a ByOrder) Swap(i,j int) {a[i],a[j]=a[j],a[i]}
func (a ByOrder) Less(i,j int) bool {return a[i].order < a[j].order}

