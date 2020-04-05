// Copyright 2018-20 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package stack

import (
	"reflect"
	"sync"
)

type tocall struct {
	fn        interface{}
	goroutine bool // Run in goroutine or not
}

// Stack is used to stores closures in order.
type Stack struct {
	lifo bool
	sync.Mutex
	stack []tocall
}

// NewStack creates a new Stack.
// lifo will make the stack operate in "LIFO" mode.
// Otherwise it will operate in "FIFO" mode.
// For fordefer, LIFO is required.
func NewStack(lifo bool, capacity ...int) *Stack {
	c := 1
	if len(capacity) > 0 {
		c = capacity[0]
	}

	return &Stack{
		lifo:  lifo,
		stack: make([]tocall, 0, c),
	}
}

// Add inserts a closure to the stack.
func (s *Stack) Add(goroutine bool, fn interface{}) {
	s.Lock()
	defer s.Unlock()

	tc := tocall{
		fn:        fn,
		goroutine: goroutine,
	}

	s.stack = append(s.stack, tc)
}

// Unwind executes all stored closures from the
// beginning to the end.
func (s *Stack) Unwind() {
	s.Lock()
	defer s.Unlock()

	if s.lifo {
		for i := len(s.stack) - 1; i >= 0; i-- {
			tc := s.stack[i]
			val := reflect.ValueOf(tc.fn)
			if tc.goroutine {
				go val.Call([]reflect.Value{})
			} else {
				val.Call([]reflect.Value{})
			}
		}
	} else {
		for i := range s.stack {
			tc := s.stack[i]
			val := reflect.ValueOf(tc.fn)
			if tc.goroutine {
				go val.Call([]reflect.Value{})
			} else {
				val.Call([]reflect.Value{})
			}
		}
	}

	// Reset stack back to original state
	s.stack = []tocall{}
}
