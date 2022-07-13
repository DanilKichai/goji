// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

import (
	"sync"
)

type ControlTerminals struct {
	Output chan interface{}
}

type Control struct {
	terminals ControlTerminals
	data      *Data
}

func NewControl(
	terminals ControlTerminals,
	data *Data,
	receiver NotifyReceiver,
) *Control {
	newControl := Control{
		terminals: terminals,
		data:      data,
	}
	receiver.RegisterControl(&newControl)

	return &newControl
}

func (my *Control) Run(parentWaitGroup *sync.WaitGroup) {
	if my.terminals.Output != nil {
		my.terminals.Output <- my.data.Get()
	}

	parentWaitGroup.Done()
}
