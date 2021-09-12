package tests

import (
	"fmt"
	"testing"

	"sabariram.com/datastructure/queue"
)

func TestQueue(t *testing.T) {
	s, err := queue.New(10)
	if err != nil {
		t.Errorf("Failed ! %v", err)
	}
	s.Enqueue("a")
	s.Enqueue(10)
	s.Enqueue(1.2)
	fmt.Printf("%+v\n", s)
	for v, err := s.Dequeue(); err == nil; v, err = s.Dequeue() {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("%+v\n", s)
	for i := 0; i < 100; i++ {
		err = s.Enqueue(i)
		if err != nil {
			fmt.Printf("%+v\n", err)
			break
		}
	}
	fmt.Printf("%+v\n", s)
	for i:= 0 ; i<5 ; i++{
		v, err := s.Dequeue()
		fmt.Printf("%v\n", v)
		if err != nil {
			fmt.Printf("%+v\n", err)
			break
		}
	}
	fmt.Printf("%+v\n", s)
	s.Extend()
	fmt.Printf("%+v\n", s)
}
