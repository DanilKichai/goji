// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

import (
	"sync"
)

type VariableTerminals struct {
	Input  chan interface{}
	Output chan interface{}
}

type Variable struct {
	terminals VariableTerminals
	data      *Data
}

func NewVariable(terminals VariableTerminals, data *Data) *Variable {
	return &Variable{
		terminals: terminals,
		data:      data,
	}
}

func (my *Variable) Run(parentWaitGroup *sync.WaitGroup) {
	if my.terminals.Input != nil {
		my.data.Set(<-my.terminals.Input, nil)
	}

	if my.terminals.Output != nil {
		my.terminals.Output <- my.data.Get()
	}

	parentWaitGroup.Done()
}
