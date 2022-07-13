// Goji.Crutch

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package crutch

import (
	"fmt"
	"os"
	"sync"
)

type PrintTerminals struct {
	Data chan any
}

type Print struct {
	terminals PrintTerminals
}

func NewPrint(terminals PrintTerminals) *Print {
	return &Print{
		terminals: terminals,
	}
}

func (my *Print) Run(parentWaitGroup *sync.WaitGroup) {
	fmt.Fprint(os.Stdout, <-my.terminals.Data)

	parentWaitGroup.Done()
}
