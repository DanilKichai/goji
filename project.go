// Goji

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package main

import (
	"goji/core"
	"goji/hello"
	"sync"
)

const (
	GlobalVariable_test1 = iota
	GlobalVariable_test2 = iota
	GlobalVariable_test3 = iota

	GlobalVariablesCount = iota
)

func initializeProjectGlobalVariables() []*core.Data {
	globalVariables := core.NewDataStore(GlobalVariablesCount)

	globalVariables[GlobalVariable_test1].Set(int(-10), nil)
	globalVariables[GlobalVariable_test2].Set(uint(10), nil)
	globalVariables[GlobalVariable_test3].Set(bool(true), nil)

	return globalVariables
}

func main() {
	waitGroup := sync.WaitGroup{}

	virtualInstruments := []core.VirtualInstrument{
		hello.NewHello(
			hello.HelloTerminals{},
			hello.InitializeHelloLocalVariables(),
			initializeProjectGlobalVariables(),
		),
	}

	waitGroup.Add(len(virtualInstruments))

	for _, vi := range virtualInstruments {
		vi.Run(&waitGroup)
	}

	waitGroup.Wait()
}
