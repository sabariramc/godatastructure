package queue

import (
	"errors"
	"fmt"
	"sync"
)

type Queue struct {
	values  []interface{}
	headPtr int
	tailPtr int
	size    int
	mu      sync.Mutex
}

func New(size int) (s *Queue, err error) {
	if size <= 0 {
		err = errors.New("size cannot be less than zero")
		return
	}
	s = &Queue{values: make([]interface{}, size), headPtr: 0, tailPtr: -1, size: 0}
	return
}

func (s *Queue) Extend() {
	s.mu.Lock()
	defer s.mu.Unlock()
	cur_len := len(s.values)
	s.values = append(s.values, make([]interface{}, cur_len)...)
	if s.headPtr > s.tailPtr && s.size > 0 {
		temp1 := s.values[0 : s.tailPtr+1]
		temp2 := s.values[cur_len : cur_len+s.tailPtr+1]
		copy(temp2, temp1)
		for i := 0; i <= s.tailPtr; i++ {
			s.values[i] = nil
		}
		s.tailPtr = cur_len + s.tailPtr
	}
}

func (s *Queue) Enqueue(value interface{}) (err error) {
	if s.size == len(s.values) {
		err = errors.New("Queue is full")
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tailPtr++
	if s.tailPtr >= len(s.values) {
		s.tailPtr = 0
	}
	s.values[s.tailPtr] = value
	s.size++
	return
}

func (s *Queue) Dequeue() (v interface{}, err error) {
	if s.size == 0 {
		err = errors.New("Queue is empty")
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	v = s.values[s.headPtr]
	s.values[s.headPtr] = nil
	s.size--
	s.headPtr++
	if s.headPtr >= len(s.values) {
		s.headPtr = 0
	}
	return
}

func (s *Queue) Peek() (v interface{}, err error) {
	if s.size == 0 {
		err = errors.New("Queue is empty")
		return
	}
	v = s.values[s.headPtr]
	return
}

func (s *Queue) String() (fmts string) {
	fmts = fmt.Sprintf("Size : %v \n", len(s.values))
	fmts += fmt.Sprintf("Head Pointer : %v \n", s.headPtr)
	fmts += fmt.Sprintf("Tail Pointer : %v \n", s.tailPtr)
	if s.size > 0 {
		fmts += fmt.Sprintf("Value List: \n")
	}
	for i := s.headPtr; ; i++ {
		if i >= len(s.values) {
			i = 0
		}
		fmts += fmt.Sprintf("%v\n", s.values[i])
		if i == s.tailPtr {
			break
		}
	}
	fmts += fmt.Sprintf("Underlying Structure : %v \n", s.values)
	return
}

func (s *Queue) Size() int {
	return len(s.values)
}
