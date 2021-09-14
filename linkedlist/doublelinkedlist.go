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
	if ll.tail == nil {
		ll.insert(&ll.head, value)
	} else {
		ll.insert(&ll.tail.next, value)
	}

}

func (ll *DoubleLinkedList) InsertAt(index int, value interface{}) (err error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	next := &ll.head
	for i := 0; i < index; i++ {
		if *next != nil {
			next = &(*next).next
		} else {
			err = errors.New("end of list reached")
			return
		}
	}
	ll.insert(next, value)
	return
}

func (ll *DoubleLinkedList) InsertAfter(searchValue interface{}, value interface{}) (err error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	next, err := ll.search(searchValue)
	if err != nil {
		return
	}
	ll.insert(&(*next).next, value)
	return
}

func (ll *DoubleLinkedList) InsertBefore(searchValue interface{}, value interface{}) (err error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	next, err := ll.search(searchValue)
	if err != nil {
		return
	}
	ll.insert(next, value)
	return
}

func (ll *DoubleLinkedList) insert(prev **node, value interface{}) {
	insert_node := newNode(value)
	temp := *prev
	*prev = insert_node
	if temp != nil {
		insert_node.next = temp
	} else {
		ll.tail = insert_node
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
	_, err = ll.search(value)
	return
}

func (ll *DoubleLinkedList) search(value interface{}) (nav **node, err error) {
	nav = &ll.head
	for *nav != nil {
		if (*nav).value == value {
			return
		}
		nav = &(*nav).next
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

func (ll *DoubleLinkedList) Size() int {
	return ll.size
}
