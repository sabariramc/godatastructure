package linkedlist

import (
	"errors"
	"fmt"
	"sync"
)

type SingleLinkedList struct {
	head *node
	tail *node
	size int
	mu   sync.Mutex
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{head: nil, tail: nil, size: 0}
}

func (ll *SingleLinkedList) Insert(value interface{}) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	if ll.tail == nil {
		ll.insert(&ll.head, value)
	} else {
		ll.insert(&ll.tail.next, value)
	}

}

func (ll *SingleLinkedList) InsertAt(index int, value interface{}) (err error) {
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

func (ll *SingleLinkedList) InsertAfter(searchValue interface{}, value interface{}) (err error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	next, err := ll.search(searchValue)
	if err != nil {
		return
	}
	ll.insert(&(*next).next, value)
	return
}

func (ll *SingleLinkedList) InsertBefore(searchValue interface{}, value interface{}) (err error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	next, err := ll.search(searchValue)
	if err != nil {
		return
	}
	ll.insert(next, value)
	return
}

func (ll *SingleLinkedList) insert(prev **node, value interface{}) {
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

func (ll *SingleLinkedList) Delete(value interface{}) (err error) {
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
			removeNode(*nav)
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

func (ll *SingleLinkedList) Search(value interface{}) (err error) {
	_, err = ll.search(value)
	return
}

func (ll *SingleLinkedList) search(value interface{}) (nav **node, err error) {
	nav = &ll.head
	for *nav != nil {
		if (*nav).value == value {
			return
		}
		nav = &(*nav).next
	}
	err = fmt.Errorf("%v - not found", value)
	return
}

func (ll *SingleLinkedList) String() (fmts string) {
	nav := ll.head
	for nav != nil {
		fmts += fmt.Sprintf("%v -> ", nav.value)
		nav = nav.next
	}
	return
}

func (ll *SingleLinkedList) Swap(a interface{}, b interface{}) (err error) {
	nodeCh := make(chan **node, 2)
	errorCh := make(chan error, 2)
	go ll.parallelSearch(a, nodeCh, errorCh)
	go ll.parallelSearch(b, nodeCh, errorCh)
	navNode := make([]**node, 2)
	idx := 0
	var temp **node
out:
	for {
		select {
		case temp = <-nodeCh:
			navNode[idx] = temp
			idx++
			if idx == 2 {
				break out
			}
		case err = <-errorCh:
			return
		}
	}
	swapNode(navNode[0], navNode[1])
	return
}

func (ll *SingleLinkedList) parallelSearch(value interface{}, nodeCh chan **node, errCh chan error) {
	nav, err := ll.search(value)
	if err != nil {
		nodeCh <- nav
	} else {
		errCh <- err
	}
}

func (ll *SingleLinkedList) Size() int {
	return ll.size
}
