// Goji.Core

// Copyright (c) 2022 Danil Kichai.

// Use, modification and distribution is subject to the Boost Software License,
// Version 1.0. (See accompanying file LICENSE_1_0.txt or copy at
// http://www.boost.org/LICENSE_1_0.txt)

package core

type Data struct {
	value       interface{}
	subscribers []DataSubscriber
}

type DataSubscriber interface {
	notify(my *Data)
}

func NewData() *Data {
	return &Data{}
}

func (my *Data) Subscribe(notify DataSubscriber) {
	my.subscribers = append(my.subscribers, notify)
}

func (my *Data) Get() interface{} {
	return my.value
}

func (my *Data) Set(value interface{}, source DataSubscriber) {
	my.value = value

	for i := range my.subscribers {
		if source == nil || source != my.subscribers[i] {
			my.subscribers[i].notify(my)
		}
	}
}

func NewDataStore(length int) []*Data {
	newDataStore := make([]*Data, length)
	for i := range newDataStore {
		newDataStore[i] = NewData()
	}

	return newDataStore
}
