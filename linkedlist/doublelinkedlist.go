package linkedlist

import (
	"errors"
	"fmt"
	"sync"
)

type DobleLinkedList struct {
	head *node
	tail *node
	size int
	mu   sync.Mutex
}

func NewDoubleLinkedList() *DobleLinkedList {
	return &DobleLinkedList{head: nil, tail: nil, size: 0}
}

func (ll *DobleLinkedList) Insert(value interface{}) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	if ll.tail == nil {
		ll.insert(&ll.head, value)
	} else {
		ll.insert(&ll.tail.next, value)
	}

}

func (ll *DobleLinkedList) InsertAt(index int, value interface{}) (err error) {
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

func (ll *DobleLinkedList) InsertAfter(searchValue interface{}, value interface{}) (err error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	next, err := ll.search(searchValue)
	if err != nil {
		return
	}
	ll.insert(&(*next).next, value)
	return
}

func (ll *DobleLinkedList) InsertBefore(searchValue interface{}, value interface{}) (err error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	next, err := ll.search(searchValue)
	if err != nil {
		return
	}
	ll.insert(next, value)
	return
}

func (ll *DobleLinkedList) insert(insertAt **node, value interface{}) {
	insertNode := newNode(value)
	temp := *insertAt
	*insertAt = insertNode
	if temp != nil {
		insertNode.next = temp
		insertNode.prev = temp.prev
		temp.prev = insertNode
	} else {
		insertNode.prev = ll.tail
		if ll.tail != nil {
			ll.tail.next = insertNode
		}
		ll.tail = insertNode
	}
	ll.size++
}

func (ll *DobleLinkedList) Delete(value interface{}) (err error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	deleteNode, err := ll.search(value)
	if err != nil {
		return
	}
	prev := (*deleteNode).prev
	next := (*deleteNode).next
	removeNode(*deleteNode)
	*deleteNode = next
	if next == nil {
		ll.tail = prev
	} else {
		next.prev = prev
	}
	if prev != nil {
		prev.next = next
	}
	return
}

func (ll *DobleLinkedList) Search(value interface{}) (err error) {
	_, err = ll.search(value)
	return
}

func (ll *DobleLinkedList) search(value interface{}) (nav **node, err error) {
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

func (ll *DobleLinkedList) String() (fmts string) {
	nav := ll.head
	if ll.head != nil {
		fmts += fmt.Sprintf("Head ptr at %v, Tail ptr at %v : ", ll.head.value, ll.tail.value)
	}
	for nav != nil {
		fmts += fmt.Sprintf("%v -> ", nav.value)
		nav = nav.next
	}
	fmts += "\nPrev Navigation:"
	nav = ll.tail
	for nav != nil {
		fmts += fmt.Sprintf("%v -> ", nav.value)
		nav = nav.prev
	}
	return
}

func (ll *DobleLinkedList) Swap(a interface{}, b interface{}) (err error) {
	ll.mu.Lock()
	defer ll.mu.Unlock()
	nodeCh := make(chan **node, 2)
	errorCh := make(chan error, 2)
	go ll.parallelSearch(a, nodeCh, errorCh)
	go ll.parallelSearch(b, nodeCh, errorCh)
	navNode := make([]**node, 2)
	var temp **node
	var isError bool
	for i := 0; i < 2; i++ {
		select {
		case temp = <-nodeCh:
			navNode[i] = temp
		case err = <-errorCh:
			isError = true
		}
	}
	if isError {
		return
	}
	if navNode[0] == navNode[1] {
		err = errors.New("cant swap same node")
		return
	}
	swapNode(navNode[0], navNode[1])
	if (*navNode[0]).next == nil {
		ll.tail = *navNode[0]
	}
	if (*navNode[1]).next == nil {
		ll.tail = *navNode[1]
	}
	return
}

func (ll *DobleLinkedList) parallelSearch(value interface{}, nodeCh chan **node, errCh chan error) {
	nav, err := ll.search(value)
	if err == nil {
		nodeCh <- nav
	} else {
		errCh <- err
	}
}

func (ll *DobleLinkedList) Size() int {
	return ll.size
}
