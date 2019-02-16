package binary_search_tree

import(
	"../../../utils"
)

type node struct {
	x 		int
	left   *node
	right  *node
	parent *node
}

func newNode() *node {
	return & node{}
}

type BinarySearchTree struct {
	r *node
}

func NewBT() *BinarySearchTree {
	return &BinarySearchTree{}
}
