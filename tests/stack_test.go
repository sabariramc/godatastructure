package tests

import (
	"fmt"
	"testing"

	"sabariram.com/datastructure/stack"
)

func TestStack(t *testing.T) {
	s, err := stack.New(10)
	//s, err = stack.NewDynamicStack(10)
	if err != nil {
		t.Errorf("Failed ! %v", err)
	}
	s.Push("a")
	s.Push(10)
	s.Push(1.2)
	fmt.Printf("%+v\n", s)
	for v, err := s.Pop(); err == nil; v, err = s.Pop() {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("%+v\n", s)
	for i:= 0; i< 100; i++ {
		s.Push(i)
	}
	for v, err := s.Pop(); err == nil; v, err = s.Pop() {
		fmt.Printf("%v\n", v)
	}
	for i:= 0; i< 10; i++ {
		s.Push(i)
	}
	fmt.Printf("%+v\n", s)
}
