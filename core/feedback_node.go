// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

import (
	"reflect"
	"sync"
)

type FeedbackNodeTerminals struct {
	Input      chan interface{}
	FirstValue chan interface{}
	Output     chan interface{}
}

type FeedbackNode struct {
	terminals     FeedbackNodeTerminals
	previousValue interface{}
	staleFlag     bool
}

func NewFeedbackNode(terminals FeedbackNodeTerminals) *FeedbackNode {
	return &FeedbackNode{
		terminals: terminals,
	}
}

func (my *FeedbackNode) Run(parentWaitGroup *sync.WaitGroup) {
	input := <-my.terminals.Input

	var firstValue interface{}
	if my.terminals.FirstValue != nil {
		firstValue = <-my.terminals.FirstValue
	} else {
		firstValue = reflect.New(reflect.TypeOf(input)).Elem().Interface()
	}

	var output interface{}
	if my.staleFlag {
		output = my.previousValue
	} else {
		output = firstValue
		my.staleFlag = true
	}

	my.previousValue = input

	if my.terminals.Output != nil {
		my.terminals.Output <- output
	}

	parentWaitGroup.Done()
}
