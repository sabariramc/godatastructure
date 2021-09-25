package tests

import (
	"fmt"
	"testing"

	"sabariram.com/datastructure/tree"
)

func compare(a *interface{}, b *interface{}) tree.CompareResult {
	intA := (*a).(int)
	intB := (*b).(int)
	if intA == intB {
		return tree.EQUAL
	} else if intA < intB {
		return tree.LESSER
	} else {
		return tree.GREATER
	}
}

func deleteNode(tr tree.Tree, index interface{}) {
	err := tr.Delete(index)
	if err == nil {
		fmt.Printf("Deleted %v; Tree: %v", index, tr)
	} else {
		fmt.Printf("Delete %v failed: %v\n", index, err)
	}
}

func TestBinarySearchTree(t *testing.T) {
	tr := tree.NewBinarySearchTree(compare)
	l := []int{10, 5, 2, 4, 1, 3, 8, 6, 7, 9, 15, 12, 13, 11, 14, 18, 16, 17, 19}
	for _, v := range l {
		tr.Insert(v, nil)
	}
	fmt.Printf("Tree: %v", tr)
	deleteNode(tr, 10)
	deleteNode(tr, 5)
	deleteNode(tr, 1)
	deleteNode(tr, 12)
	deleteNode(tr, 15)

}
