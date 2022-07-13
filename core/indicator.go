// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

import (
	"sync"
)

type IndicatorTerminals struct {
	Input chan interface{}
}

type Indicator struct {
	terminals IndicatorTerminals
	data      *Data
}

func NewIndicator(
	terminals IndicatorTerminals,
	data *Data,
	receiver NotifyReceiver,
) *Indicator {
	newIndicator := Indicator{
		terminals: terminals,
		data:      data,
	}
	receiver.RegisterIndicator(&newIndicator)

	return &newIndicator
}

func (my *Indicator) Run(parentWaitGroup *sync.WaitGroup) {
	if my.terminals.Input != nil {
		my.data.Set(<-my.terminals.Input, nil)
	}

	parentWaitGroup.Done()
}
