package linkedlist

type node struct {
	value interface{}
	next  *node
	prev  *node
}

func newNode(a interface{}) *node {
	return &node{value: a, next: nil, prev: nil}
}

func swapNode(a **node, b **node) {
	aNext, aPrev := a.next, a.prev
	a.next, a.prev = b.next, b.prev
	b.next, b.prev = aNext, aPrev
}

func removeNode(a *node) {
	a.next = nil
	a.prev = nil
}
