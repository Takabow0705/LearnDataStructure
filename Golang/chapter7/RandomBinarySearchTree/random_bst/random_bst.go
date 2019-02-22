package random_bst

import(
	"../../../utils"
)

type struct node{
	x int
	left *node
	right *node
	parent *node
	p int //priority
}

type struct Treap{
	r *node
	n int
}
