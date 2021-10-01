package tree

import (
	"errors"
	"fmt"
)

type AVLTree struct {
	root          *node
	size          int
	compare       CompareFunction
	rotationCount int
}

func NewAVLTree(compare CompareFunction) *AVLTree {
	return &AVLTree{root: nil, size: 0, compare: compare}
}

func (bt *AVLTree) Insert(index interface{}, value *interface{}) error {
	if index == nil {
		return errors.New("index cannot be nil")
	}
	nav := bt.search(&bt.root, index)
	if *nav != nil {
		return errors.New("duplicate index")
	}
	bt.rotationCount = 0
	_, err := bt.insert(&bt.root, &index, value)
	fmt.Printf("No of rotations - %v \n", bt.rotationCount)
	return err
}

func (bt *AVLTree) GetRotationCount() int {
	return bt.rotationCount
}

func (bt *AVLTree) Search(index interface{}) (value *interface{}, err error) {
	nav := bt.search(&bt.root, index)
	if nav != nil {
		value = &((*nav).value)
	} else {
		err = errors.New("index not found")
	}
	return
}

func (bt *AVLTree) Replace(index interface{}, value *interface{}) (err error) {
	nav := bt.search(&bt.root, index)
	if nav != nil {
		(*nav).value = *value
	} else {
		err = errors.New("index not found")
	}
	return
}

func (bt *AVLTree) insert(nav **node, index *interface{}, value *interface{}) (height int, err error) {
	if *nav == nil {
		*nav = newNode(*index, value)
		bt.size++
		height = 1
		return
	}
	result := bt.compare(&(*nav).index, index)
	switch result {
	case EQUAL:
		err = errors.New("duplicate index")
	case LESSER:
		(*nav).rightHeight, err = bt.insert(&(*nav).right, index, value)
	case GREATER:
		(*nav).leftHeight, err = bt.insert(&(*nav).left, index, value)
	}
	if err != nil {
		return
	}
	bt.balanceNode(nav)
	height = bt.getHeight(*nav)
	return

}

func (bt *AVLTree) getHeight(temp *node) int {
	if temp == nil {
		return 0
	}
	lh, rh := temp.leftHeight, temp.rightHeight
	if lh > rh {
		return lh + 1
	} else {
		return rh + 1
	}
}

func (bt *AVLTree) getBalanceFactor(temp *node) int {
	if temp == nil {
		return 0
	}
	return temp.rightHeight - temp.leftHeight
}

func (bt *AVLTree) search(nav **node, index interface{}) **node {
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

func (bt *AVLTree) Delete(index interface{}) (err error) {
	if index == nil {
		return errors.New("index cannot be nil")
	}
	nav := bt.search(&bt.root, index)
	if *nav != nil {
		return errors.New("duplicate index")
	}
	bt.rotationCount = 0
	_, err = bt.delete(&bt.root, &index)
	fmt.Printf("No of rotations - %v \n", bt.rotationCount)
	return err
}

func (bt *AVLTree) delete(nav **node, index *interface{}) (height int, err error) {
	if *nav == nil {
		err = errors.New("index not found")
		return
	}
	result := bt.compare(&(*nav).index, index)
	switch result {
	case EQUAL:
		temp := *nav
		*nav = bt.getSuccessor(*nav)
		if *nav != nil {
			(*nav).left = temp.left
			(*nav).right = temp.right
		}
		removeNode(temp)
		bt.size--
	case LESSER:
		(*nav).rightHeight, err = bt.delete(&(*nav).right, index)
	case GREATER:
		(*nav).leftHeight, err = bt.delete(&(*nav).left, index)
	}
	if err != nil {
		return
	}
	bt.balanceNode(nav)
	height = bt.getHeight(*nav)
	return
}

func (bt *AVLTree) getSuccessor(temp *node) *node {
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
		temp.right = nil
		return right
	} else {
		temp.left = nil
		return left
	}
}

func (bt *AVLTree) String() (fmts string) {
	fmts = fmt.Sprintf("BST size: %v\n", bt.size)
	fmts += fmt.Sprintf("In order index: %v\n", getIndexLine(inOrderTraversal, bt.root, bt.size))
	fmts += fmt.Sprintf("Pre order index: %v\n", getIndexLine(preOrderTraversal, bt.root, bt.size))
	fmts += fmt.Sprintf("Post order index: %v\n", getIndexLine(postOrderTraversal, bt.root, bt.size))
	return
}

func (bt *AVLTree) Validate() error {
	return nil
}

func (bt *AVLTree) balanceNode(nav **node) {
	if *nav == nil {
		return
	}
	balanceFactor := bt.getBalanceFactor(*nav)
	if balanceFactor < -1 {
		nav = &(*nav).left
		balanceFactor = bt.getBalanceFactor(*nav)
		if balanceFactor <= 0 {
			bt.rotateRight(nav)
		} else {
			bt.rotateRightLeft(nav)
		}
	} else if balanceFactor > 1 {
		nav = &(*nav).right
		balanceFactor = bt.getBalanceFactor(*nav)
		if balanceFactor >= 0 {
			bt.rotateLeft(nav)
		} else {
			bt.rotateLeftRight(nav)
		}
	}
}

func (bt *AVLTree) rotateRight(nav **node) {
	target := *nav
	left := target.left
	left_right := left.right
	*nav = left
	left.right = target
	target.left = left_right
	target.leftHeight = bt.getHeight(target.left)
	left.rightHeight = bt.getHeight(target)
	bt.rotationCount++
}

func (bt *AVLTree) rotateLeft(nav **node) {
	target := *nav
	right := target.right
	right_left := right.left
	*nav = right
	right.left = target
	target.right = right_left
	target.rightHeight = bt.getHeight(target.right)
	right.leftHeight = bt.getHeight(right.left)
	bt.rotationCount++
}

func (bt *AVLTree) rotateRightLeft(nav **node) {
	target := (*nav).left
	right := target.right
	right_left := right.left
	target.right = right_left
	right.left = target
	(*nav).left = right
	target.rightHeight = bt.getHeight(target.right)
	right.leftHeight = bt.getHeight(right.left)
	(*nav).leftHeight = bt.getHeight(right)
	bt.rotationCount++
	bt.rotateRight(nav)
}

func (bt *AVLTree) rotateLeftRight(nav **node) {
	target := (*nav).right
	left := target.left
	left_right := left.right
	target.left = left_right
	left.right = target
	(*nav).right = left
	target.leftHeight = bt.getHeight(target.left)
	left.rightHeight = bt.getHeight(left.right)
	(*nav).rightHeight = bt.getHeight(left)
	bt.rotationCount++
	bt.rotateLeft(nav)
}
