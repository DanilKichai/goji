// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

import (
	"sync"
)

type JunctionTerminals struct {
	Input   chan interface{}
	Outputs []chan interface{}
}

type Junction struct {
	terminals JunctionTerminals
}

func NewJunction(terminals JunctionTerminals) *Junction {
	return &Junction{
		terminals: terminals,
	}
}

func (my *Junction) Run(parentWaitGroup *sync.WaitGroup) {
	input := <-my.terminals.Input

	for i := range my.terminals.Outputs {
		my.terminals.Outputs[i] <- input
	}

	parentWaitGroup.Done()
}
