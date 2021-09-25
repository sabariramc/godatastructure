package tree

type CompareFunction func(a *interface{}, b *interface{}) CompareResult

type Tree interface {
	Insert(index interface{}, value *interface{}) error
	Delete(index interface{}) error
	Search(index interface{}) (*interface{}, error)
	Validate() error
	Replace(index interface{}, value *interface{}) error
}
