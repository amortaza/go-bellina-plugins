package hover

func On(cb func(interface{})) {
	plugin.On(cb)
}
