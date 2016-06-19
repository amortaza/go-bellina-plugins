package zindex

type Event struct {
}

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "zindex"
}

func (c *Plugin) Init() {
	g_ctxByNodeId = make(map[string] *Ctx)
}

func (c *Plugin) Tick() {
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not supoorted in click.Plugin")
}

func NewPlugin() *Plugin {
	c := &Plugin{}

	return c
}
