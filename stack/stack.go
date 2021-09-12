package stack

import (
	"errors"
	"fmt"
	"sync"
)

type Stack struct {
	values []interface{}
	ptr    int
	mu sync.Mutex
}

func New(size int) (s *Stack, err error) {
	if size <= 0 {
		err = errors.New("size cannot be less than zero")
		return
	}
	s = &Stack{}
	s.values = make([]interface{}, size)
	s.ptr = 0
	return
}

func (s *Stack) Extend() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values = append(s.values, make([]interface{}, len(s.values))...)
}

func (s *Stack) Push(value interface{}) (err error) {
	if len(s.values) == s.ptr {
		err = errors.New("Stack is full")
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[s.ptr] = value
	s.ptr++
	return
}

func (s *Stack) Pop() (v interface{}, err error) {
	if s.ptr == 0 {
		err = errors.New("Stack is empty")
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ptr--
	v = s.values[s.ptr]
	s.values[s.ptr] = nil
	return
}

func (s *Stack) Peek() (v interface{}, err error) {
	if s.ptr == 0 {
		err = errors.New("Stack is empty")
		return
	}
	v = s.values[s.ptr-1]
	return
}

func (s *Stack) String() (fmts string) {
	fmts = fmt.Sprintf("Size : %v \n", len(s.values))
	fmts += fmt.Sprintf("Stack Pointer : %v \n", s.ptr)
	if s.ptr > 0 {
		fmts += fmt.Sprintf("Value List: \n")
	}
	for i := 0; i < s.ptr; i++ {
		fmts += fmt.Sprintf("%v\n", s.values[i])
	}
	return
}

func (s *Stack) Size() int {
	return len(s.values)
}
