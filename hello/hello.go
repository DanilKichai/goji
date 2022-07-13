// Goji.Hello

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package hello

import (
	"goji/core"
	"goji/crutch"
)

const (
	localVariable_test1 = iota
	localVariable_test2 = iota
	localVariable_test3 = iota

	localVariablesCount = iota
)

func InitializeHelloLocalVariables() []*core.Data {
	localVariables := core.NewDataStore(localVariablesCount)

	localVariables[localVariable_test1].Set(int(32), nil)
	localVariables[localVariable_test2].Set(bool(true), nil)
	localVariables[localVariable_test3].Set(string("test"), nil)

	return localVariables
}

type HelloTerminals struct {
}

func NewHello(
	terminals HelloTerminals,
	localVariables []*core.Data,
	globalVariables []*core.Data,
) *core.SubVirtualInstrument {
	wires := core.NewWireHarness(1)

	return core.NewSubVirtualInstrument(
		core.SubVirtualInstrumentTerminals{},
		localVariables,
		globalVariables,
		[]core.VirtualInstrument{
			core.NewConstant(
				core.ConstantTerminals{
					Output: wires[0],
				},
				"Hello, world!\n",
			),
			crutch.NewPrint(
				crutch.PrintTerminals{
					Data: wires[0],
				},
			),
		},
	)
}
