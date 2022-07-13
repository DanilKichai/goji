// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

func NewWire() chan interface{} {
	return make(chan interface{})
}

func NewWireHarness(length int) []chan interface{} {
	newWireHarness := make([]chan interface{}, length)
	for i := range newWireHarness {
		newWireHarness[i] = NewWire()
	}
	return newWireHarness
}
