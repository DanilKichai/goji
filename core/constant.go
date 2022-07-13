// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

import (
	"sync"
)

type ConstantTerminals struct {
	Output chan interface{}
}

type Constant struct {
	terminals ConstantTerminals
	value     interface{}
}

func NewConstant(terminals ConstantTerminals, value interface{}) *Constant {
	return &Constant{
		terminals: terminals,
		value:     value,
	}
}

func (my *Constant) Run(parentWaitGroup *sync.WaitGroup) {
	if my.terminals.Output != nil {
		my.terminals.Output <- my.value
	}

	parentWaitGroup.Done()
}
