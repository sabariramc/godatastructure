package tests

import (
	"fmt"
	"testing"

	"sabariram.com/datastructure/linkedlist"
)

func deleteOperation(ll linkedlist.LinkedList, value interface{}) {
	err := ll.Delete(value)
	if err != nil {
		fmt.Printf("Delete %v failed: %v \n", value, err)
		return
	}
	fmt.Printf("Deleted %v : %v \n", value, ll)
}

func insertOperation(ll linkedlist.LinkedList, value interface{}) {
	ll.Insert(value)
	fmt.Printf("Inserted %v : %v \n", value, ll)
}

func insertAtOperation(ll linkedlist.LinkedList, value interface{}, index int) {
	err := ll.InsertAt(index, value)
	if err != nil {
		fmt.Printf("Insert %v At %v Failed : %v \n", value, index, err)
	} else {
		fmt.Printf("Inserted %v At %v : %v \n", value, index, ll)
	}
}

func insertAfterOperation(ll linkedlist.LinkedList, value interface{}, searchValue interface{}) {
	err := ll.InsertAfter(searchValue, value)
	if err != nil {
		fmt.Printf("Insert %v After %v Failed : %v \n", value, searchValue, err)
	} else {
		fmt.Printf("Inserted %v After %v : %v \n", value, searchValue, ll)
	}
}

func insertBeforeOperation(ll linkedlist.LinkedList, value interface{}, searchValue interface{}) {
	err := ll.InsertBefore(searchValue, value)
	if err != nil {
		fmt.Printf("Insert %v Before %v Failed : %v \n", value, searchValue, err)
	} else {
		fmt.Printf("Inserted %v Before %v : %v \n", value, searchValue, ll)
	}
}

func searchOperataion(ll linkedlist.LinkedList, value interface{}) {
	err := ll.Search(value)
	if err != nil {
		fmt.Printf("Searched %v : %v \n", value, err)
	} else {
		fmt.Printf("Searched %v : found \n", value)
	}
}

func swapOperation(ll linkedlist.LinkedList, valueA interface{}, valueB interface{}) {
	err := ll.Swap(valueA, valueB)
	if err != nil {
		fmt.Printf("Swap %v, %v Failed: %v \n", valueA, valueB, err)
	} else {
		fmt.Printf("Swaped %v, %v : %v \n", valueA, valueB, ll)
	}
}

func TestSinglyLinkedList(t *testing.T) {
	ll := linkedlist.NewSingleLinkedList()
	deleteOperation(ll, 123)
	searchOperataion(ll, 1)
	insertOperation(ll, 1)
	insertOperation(ll, 2)
	insertOperation(ll, 3)
	searchOperataion(ll, 1)
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
	searchOperataion(ll, 3)
	searchOperataion(ll, 4)
	searchOperataion(ll, 9)
	searchOperataion(ll, 8)
	insertAtOperation(ll, 1, 0)
	insertAtOperation(ll, 7, 3)
	insertAtOperation(ll, 10, 6)
	insertAtOperation(ll, 11, 10)
	insertAfterOperation(ll, 20, 10)
	insertBeforeOperation(ll, -1, 1)
	insertBeforeOperation(ll, 0, 1)
	insertAfterOperation(ll, 11, 10)
	insertBeforeOperation(ll, 0, -2)
	insertAfterOperation(ll, 11, 100)
	swapOperation(ll, 1, 4)
	swapOperation(ll, 1, 4)
	swapOperation(ll, -1, 20)
	swapOperation(ll, -1, 20)
	swapOperation(ll, 0, 5)
	swapOperation(ll, 0, 5)
}

func TestDoubleLinkedList(t *testing.T) {
	ll := linkedlist.NewDoubleLinkedList()
	deleteOperation(ll, 123)
	searchOperataion(ll, 1)
	insertOperation(ll, 1)
	insertOperation(ll, 2)
	insertOperation(ll, 3)
	searchOperataion(ll, 1)
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
	searchOperataion(ll, 3)
	searchOperataion(ll, 4)
	searchOperataion(ll, 9)
	searchOperataion(ll, 8)
	insertAtOperation(ll, 1, 0)
	insertAtOperation(ll, 7, 3)
	insertAtOperation(ll, 10, 6)
	insertAtOperation(ll, 11, 10)
	insertAfterOperation(ll, 20, 10)
	insertBeforeOperation(ll, -1, 1)
	insertBeforeOperation(ll, 0, 1)
	insertAfterOperation(ll, 11, 10)
	insertBeforeOperation(ll, 0, -2)
	insertAfterOperation(ll, 11, 100)
	swapOperation(ll, 1, 4)
	swapOperation(ll, 1, 4)
	swapOperation(ll, -1, 20)
	swapOperation(ll, -1, 20)
	swapOperation(ll, 0, 5)
	swapOperation(ll, 0, 5)
}
