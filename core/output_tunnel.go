// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

import "sync"

const (
	OutputTunnelModeLastValue = iota
	OutputTunnelModeIndexing
	OutputTunnelModeConcentrating
	OutputTunnelModeConditional
)

type OutputTunnelTerminals struct {
	Inside  chan interface{}
	Outside chan interface{}
}

type OutputTunnel struct {
	terminlas InputTunnelTerminals
	mode      int
}

func NewOutputTunnel(terminals InputTunnelTerminals, mode int) *OutputTunnel {
	return &OutputTunnel{
		terminlas: terminals,
		mode:      mode,
	}
}

func (my *OutputTunnel) ExternalEntry(parentWaitGroup *sync.WaitGroup) {
	//TODO

	parentWaitGroup.Done()
}

func (my *OutputTunnel) InternalEntry(parentWaitGroup *sync.WaitGroup) {
	//TODO

	parentWaitGroup.Done()
}
