package resize

import "github.com/amortaza/go-bellina"

func Use() {
	plugin.On(nil)
}

func On(cb func(interface{})) {
	plugin.On(cb)
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}

