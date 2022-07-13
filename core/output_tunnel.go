// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

const (
	OutputTunnelLoopModeLastValue = iota
	OutputTunnelLoopModeIndexing
	OutputTunnelLoopModeConcentrating
	OutputTunnelLoopModeConditional
)

type OutputTunnel struct {
	Inside   chan interface{}
	Outside  chan interface{}
	LoopMode int
}
