package linkedlist

import (
	"errors"
	"fmt"
	"sync"
)

type DoubleLinkedList struct {
	head *node
	tail *node
	size int
	mu   sync.Mutex
}

func NewDobuleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{head: nil, tail: nil, size: 0}
}

func (ll *DoubleLinkedList) Insert(value interface{}) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	temp := newNode(value)
	if ll.tail == nil {
		ll.head = temp
		ll.tail = temp
	} else {
		ll.tail.next = temp
		ll.tail = temp
	}
	ll.size++
}

func (ll *DoubleLinkedList) Delete(value interface{}) (err error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	nav := &ll.head
	var prev *node
	for *nav != nil {
		if (*nav).value == value {
			temp := (*nav).next
			if temp == nil {
				ll.tail = prev
			}
			(*nav).next = nil
			*nav = temp
			ll.size--
			return
		}
		prev = *nav
		nav = &(*nav).next
	}
	err = errors.New("value not found")
	return
}

func (ll *DoubleLinkedList) Search(value interface{}) (err error) {
	nav := ll.head
	for nav != nil {
		if nav.value == value {
			return
		}
		nav = nav.next
	}
	err = errors.New("value not found")
	return
}

func (ll *DoubleLinkedList) String() (fmts string) {
	nav := ll.head
	for nav != nil {
		fmts += fmt.Sprintf("%v -> ", nav.value)
		nav = nav.next
	}
	return
}
