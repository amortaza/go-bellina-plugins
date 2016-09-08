package sideglue

import (
	"fmt"
)

func fake() {
    var _ = fmt.Println
}

func Id() *State {
	return ensureState()
}


