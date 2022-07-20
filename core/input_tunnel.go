// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

import "sync"

type InputTunnelTerminals struct {
	Inside  chan interface{}
	Outside chan interface{}
}

type InputTunnel struct {
	terminlas InputTunnelTerminals
	indexing  bool
}

func NewInputTunnel(terminals InputTunnelTerminals, indexing bool) *InputTunnel {
	return &InputTunnel{
		terminlas: terminals,
		indexing:  indexing,
	}
}

func (my *InputTunnel) ExternalEntry(parentWaitGroup *sync.WaitGroup) {
	//TODO

	parentWaitGroup.Done()
}

func (my *InputTunnel) InternalEntry(parentWaitGroup *sync.WaitGroup) {
	//TODO

	parentWaitGroup.Done()
}
