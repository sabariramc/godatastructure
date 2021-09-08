package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	values  []interface{}
	headPtr int
	tailPtr int
	size    int
}

func New(size int) (s *Queue, err error) {
	if size <= 0 {
		err = errors.New("size cannot be less than zero")
		return
	}
	s = &Queue{values: make([]interface{}, size), headPtr: 0, tailPtr: 0, size: 0}
	return
}

func (s *Queue) Extend() {
	s.values = append(s.values, make([]interface{}, len(s.values)))
}

func (s *Queue) Enqueue(value interface{}) (err error) {
	if s.headPtr == s.tailPtr {
		err = errors.New("Queue is full")
		return
	}
	s.values[s.ptr] = value
	s.ptr++
	return
}

func (s *Queue) Pop() (v interface{}, err error) {
	if s.ptr == 0 {
		err = errors.New("Queue is empty")
		return
	}
	s.ptr--
	v = s.values[s.ptr]
	s.values[s.ptr] = nil
	return
}

func (s *Queue) Peek() (v interface{}, err error) {
	if s.ptr == 0 {
		err = errors.New("Queue is empty")
		return
	}
	v = s.values[s.ptr-1]
	return
}

func (s *Queue) String() (fmts string) {
	fmts = fmt.Sprintf("Size : %v \n", len(s.values))
	fmts += fmt.Sprintf("Queue Pointer : %v \n", s.ptr)
	if s.ptr > 0 {
		fmts += fmt.Sprintf("Value List: \n")
	}
	for i := 0; i < s.ptr; i++ {
		fmts += fmt.Sprintf("%v\n", s.values[i])
	}
	return
}

func (s *Queue) Size() int {
	return len(s.values)
}
