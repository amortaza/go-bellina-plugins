package hover

import "github.com/amortaza/go-bellina"

func On(cb func(interface{})) {
	plugin.On(cb)
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}
