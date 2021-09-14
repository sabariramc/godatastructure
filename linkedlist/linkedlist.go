package linkedlist

type LinkedList interface {
	Insert(value interface{})
	InsertAt(index int, value interface{}) (err error)
	InsertAfter(searchValue interface{}, value interface{}) (err error)
	InsertBefore(searchValue interface{}, value interface{}) (err error)
	Delete(value interface{}) (err error)
	Search(value interface{}) (err error)
	String() (fmts string)
	Size() int
	Swap(a interface{}, b interface{}) (err error)
}
