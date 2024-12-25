package internal

type GraphEdge struct {
	To     int
	Weight int
}

type WeightedAdjacencyList [][]GraphEdge
type WeightedAdjacencyMatrix = [][]int

type BinaryNode[T any] struct {
	Value       T
	Left, Right *BinaryNode[T]
}

type Point struct {
	X int
	Y int
}
