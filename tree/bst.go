package tree

import (
	"errors"
	"fmt"
)

type BinarySearchTree struct {
	root    *node
	size    int
	compare CompareFunction
}

func NewBinarySearchTree(compare CompareFunction) *BinarySearchTree {
	return &BinarySearchTree{root: nil, size: 0, compare: compare}
}

func (bt *BinarySearchTree) Insert(index interface{}, value *interface{}) error {
	if index == nil {
		return errors.New("index cannot be nil")
	}
	nav := bt.search(&bt.root, index)
	if *nav != nil {
		return errors.New("duplicate index")
	}
	bt.insert(nav, index, value)
	return nil
}

func (bt *BinarySearchTree) Search(index interface{}) (value *interface{}, err error) {
	nav := bt.search(&bt.root, index)
	if nav != nil {
		value = &((*nav).value)
	} else {
		err = errors.New("index not found")
	}
	return
}

func (bt *BinarySearchTree) Replace(index interface{}, value *interface{}) (err error) {
	nav := bt.search(&bt.root, index)
	if nav != nil {
		(*nav).value = *value
	} else {
		err = errors.New("index not found")
	}
	return
}

func (bt *BinarySearchTree) insert(nav **node, index interface{}, value *interface{}) {
	*nav = newNode(index, value)
	bt.size++
}

func (bt *BinarySearchTree) search(nav **node, index interface{}) **node {
	for *nav != nil {
		result := bt.compare(&(*nav).index, &index)
		switch result {
		case EQUAL:
			return nav
		case LESSER:
			nav = &(*nav).right
		case GREATER:
			nav = &(*nav).left
		}
	}
	return nav
}

func (bt *BinarySearchTree) Delete(index interface{}) (err error) {
	nav := bt.search(&bt.root, index)
	if *nav != nil {
		temp := *nav
		*nav = bt.getSuccessor(*nav)
		if *nav != nil {
			(*nav).left = temp.left
			(*nav).right = temp.right
		}
		removeNode(temp)
		bt.size--
	} else {
		err = errors.New("index not found")
	}
	return
}

func (bt *BinarySearchTree) getSuccessor(temp *node) *node {
	left := temp.left
	right := temp.right
	if left != nil && right != nil {
		nav := &temp.left
		for (*nav).right != nil {
			nav = &(*nav).right
		}
		temp = *nav
		*nav = temp.left
		return temp
	} else if left == nil {
		return right
	} else {
		return left
	}
}

func (bt *BinarySearchTree) String() (fmts string) {
	fmts = fmt.Sprintf("BST size: %v\n", bt.size)
	fmts += fmt.Sprintf("In order index: %v\n", getIndexLine(inOrderTraversal, bt.root, bt.size))
	fmts += fmt.Sprintf("Pre order index: %v\n", getIndexLine(preOrderTraversal, bt.root, bt.size))
	fmts += fmt.Sprintf("Post order index: %v\n", getIndexLine(postOrderTraversal, bt.root, bt.size))
	return
}

func getIndexLine(traversal traversalFunction, nav *node, size int) string {
	ch := make(chan *interface{}, size)
	defer close(ch)
	traversal(nav, ch)
	traverseResponse := ""
	for i := 0; i < size; i++ {
		j := <-ch
		traverseResponse += fmt.Sprintf("%v -> ", *j)
	}
	return traverseResponse
}

func (bt *BinarySearchTree) Validate() error {
	return nil
}
