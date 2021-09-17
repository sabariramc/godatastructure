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
	nodeA := *a
	nodeB := *b
	if nodeA.next == nodeB {
		nodeA.next = nodeB.next
		nodeB.next = nodeA
		if nodeB.prev != nil {
			nodeB.prev = nodeA.prev
			nodeA.prev = nodeB
		}
		*a = nodeB
	} else if nodeB.next == nodeA {
		nodeB.next = nodeA.next
		nodeA.next = nodeB
		if nodeA.prev != nil {
			nodeA.prev = nodeB.prev
			nodeB.prev = nodeA
		}
		*b = nodeA
	} else {
		aNext, aPrev := nodeA.next, nodeA.prev
		nodeA.next, nodeA.prev = nodeB.next, nodeB.prev
		nodeB.next, nodeB.prev = aNext, aPrev
		*b = nodeA
		*a = nodeB
	}
	if nodeA.next != nil && nodeA.next.prev != nil {
		nodeA.next.prev = nodeA
	}
	if nodeB.next != nil && nodeB.next.prev != nil {
		nodeB.next.prev = nodeB
	}

}

func removeNode(a *node) {
	a.value = nil
	a.next = nil
	a.prev = nil
}
