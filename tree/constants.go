package tree

type colour bool

const (
	RED   colour = true
	BLACK colour = false
)

type CompareResult int

const (
	LESSER CompareResult = -1
	EQUAL CompareResult = 0
	GREATER CompareResult = 1
)
