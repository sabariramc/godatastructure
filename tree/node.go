package tree

type node struct {
	index         interface{}
	value         interface{}
	left          *node
	right         *node
	colour        colour
	balanceFactor int
	leftHeight    int
	rightHeight   int
}

func newNode(index interface{}, value *interface{}) *node {
	return &node{value: value, index: index}
}

func removeNode(n *node) {
	n.index = nil
	n.value = nil
	n.left = nil
	n.right = nil
}

type traversalFunction func(nav *node, ch chan *interface{})

func inOrderTraversal(nav *node, ch chan *interface{}) {
	if nav == nil {
		return
	}
	inOrderTraversal(nav.left, ch)
	ch <- &nav.index
	inOrderTraversal(nav.right, ch)
}

func preOrderTraversal(nav *node, ch chan *interface{}) {
	if nav == nil {
		return
	}
	ch <- &nav.index
	preOrderTraversal(nav.left, ch)
	preOrderTraversal(nav.right, ch)
}

func postOrderTraversal(nav *node, ch chan *interface{}) {
	if nav == nil {
		return
	}
	postOrderTraversal(nav.left, ch)
	postOrderTraversal(nav.right, ch)
	ch <- &nav.index
}
