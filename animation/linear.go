package animation

func Linear(state *AnimState) (func() (float32, bool)) {

	return func() (float32, bool) {
		pct, valid := state.nextPct()

		return state.StartValue + state.diff * pct, valid
	}
}


