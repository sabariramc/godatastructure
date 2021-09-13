package tests

import (
	"fmt"
	"testing"

	"sabariram.com/datastructure/linkedlist"
)

func deleteOperation(ll *linkedlist.SingleLinkedList, value interface{}) {
	err := ll.Delete(value)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Deleted %v : %v \n", value, ll)
}

func insertOperation(ll *linkedlist.SingleLinkedList, value interface{}) {
	ll.Insert(value)
	fmt.Printf("Inserted %v : %v \n", value, ll)
}

func TestSinglyLinkedList(t *testing.T) {
	ll := linkedlist.NewSingleLinkedList()
	deleteOperation(ll, 123)
	insertOperation(ll, 1)
	insertOperation(ll, 2)
	insertOperation(ll, 3)
	deleteOperation(ll, 1)
	insertOperation(ll, 4)
	insertOperation(ll, 5)
	insertOperation(ll, 6)
	deleteOperation(ll, 3)
	deleteOperation(ll, 2)
	deleteOperation(ll, 6)
	insertOperation(ll, 7)
	insertOperation(ll, 8)
	insertOperation(ll, 9)
	deleteOperation(ll, 7)
}
