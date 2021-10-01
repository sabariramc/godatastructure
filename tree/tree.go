package tree

import "fmt"

type CompareFunction func(a *interface{}, b *interface{}) CompareResult

type Tree interface {
	Insert(index interface{}, value *interface{}) error
	Delete(index interface{}) error
	Search(index interface{}) (*interface{}, error)
	Validate() error
	Replace(index interface{}, value *interface{}) error
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