package internal

type FnSchema map[string]Func

type Func struct {
	Args   string
	Return string
}

var DEFAULT_RETURN = map[string]string{
	"bool":                 "false",
	"int":                  "0",
	"string":               "\"\"",
	"[]int":                "nil",
	"[]string":             "nil",
	"[]bool":               "nil",
	"[]in.Point":           "nil",
	"[]in.GraphEdge":       "nil",
	"[]in.BinaryNode[int]": "nil",
}

var FnS = FnSchema{
	"BubbleSort": Func{
		Args:   "arr []int",
		Return: "",
	},
	"InsertionSort": Func{
		Args:   "arr []int",
		Return: "",
	},
	"MergeSort": Func{
		Args:   "arr []int",
		Return: "",
	},
	"QuickSort": Func{
		Args:   "arr []int",
		Return: "",
	},
	"DijkstraList": Func{
		Args:   "source int, sink int, arr in.WeightedAdjacencyList",
		Return: "[]int",
	},
	"BinarySearchList": Func{
		Args:   "haystack []int, needle int",
		Return: "bool",
	},
	"LinearSearchList": Func{
		Args:   "haystack []int, needle int",
		Return: "bool",
	},
	"TwoCrystalBalls": Func{
		Args:   "breaks []bool",
		Return: "int",
	},
	"MazeSolver": Func{
		Args:   "maze []string, wall string, start in.Point, end in.Point",
		Return: "[]in.Point",
	},
	"BTPreOrder": Func{
		Args:   "head *in.BinaryNode[int]",
		Return: "[]int",
	},
	"BTInOrder": Func{
		Args:   "head *in.BinaryNode[int]",
		Return: "[]int",
	},
	"BTPostOrder": Func{
		Args:   "head *in.BinaryNode[int]",
		Return: "[]int",
	},
	"BTBFS": Func{
		Args:   "head *in.BinaryNode[int], needle int",
		Return: "bool",
	},
	"CompareBinaryTrees": Func{
		Args:   "a *in.BinaryNode[int], b *in.BinaryNode[int]",
		Return: "bool",
	},
	"DFSOnBST": Func{
		Args:   "head *in.BinaryNode[int], needle int",
		Return: "bool",
	},
	"DFSGraphList": Func{
		Args:   "graph in.WeightedAdjacencyList, source int, needle int",
		Return: "[]int",
	},
	"BFSGraphList": Func{
		Args:   "graph in.WeightedAdjacencyList, source int, needle int",
		Return: "[]int",
	},
	"BFSGraphMatrix": Func{
		Args:   "graph in.WeightedAdjacencyMatrix, source int, needle int",
		Return: "[]int",
	},
}
