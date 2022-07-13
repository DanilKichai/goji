// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

import "sync"

type VirtualInstrument interface {
	Run(parentWaitGroup *sync.WaitGroup)
}

type SubVirtualInstrumentTerminals struct {
	Inputs  []chan interface{}
	Outputs []chan interface{}
}

type SubVirtualInstrument struct {
	terminals          SubVirtualInstrumentTerminals
	localVariables     []*Data
	globalVariables    []*Data
	virtualInstruments []VirtualInstrument
	waitGroup          sync.WaitGroup
}

func NewSubVirtualInstrument(
	terminals SubVirtualInstrumentTerminals,
	localVariables []*Data,
	globalVariables []*Data,
	virtualInstruments []VirtualInstrument,
) *SubVirtualInstrument {
	return &SubVirtualInstrument{
		terminals:          terminals,
		localVariables:     localVariables,
		globalVariables:    globalVariables,
		virtualInstruments: virtualInstruments,
	}
}

func (my *SubVirtualInstrument) Run(parentWaitGroup *sync.WaitGroup) {
	my.waitGroup.Add(len(my.virtualInstruments))

	for _, subVirtualInstrument := range my.virtualInstruments {
		go subVirtualInstrument.Run(&my.waitGroup)
	}

	my.waitGroup.Wait()

	parentWaitGroup.Done()
}
