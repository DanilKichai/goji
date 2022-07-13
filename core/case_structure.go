// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

import "sync"

type CaseStructureTerminals struct {
	InputTunnels  []InputTunnel
	OutputTunnels []OutputTunnel
}

type CaseCouple struct {
	Value                interface{}
	SubVirtualInstrument VirtualInstrument
}

type CaseStructere struct {
	terminals   CaseStructureTerminals
	caseCouples []CaseCouple
	defaultCase int
}

func NewCaseStructure(
	terminals CaseStructureTerminals,
	caseCouples []CaseCouple,
	defaultCase int,
) *CaseStructere {
	return &CaseStructere{
		terminals:   terminals,
		caseCouples: caseCouples,
		defaultCase: defaultCase,
	}
}

func (my *CaseStructere) Run(parentWaitGroup *sync.WaitGroup) {
	// TODO

	parentWaitGroup.Done()
}
