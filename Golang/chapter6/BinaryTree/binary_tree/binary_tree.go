package binary_tree

type node struct{
	left *node
	right *node
	parent *node
}

func newNode() *node{
	retrun &node{}
}

type BinaryTree struct{
	r *node
}

func NewBTNOode() *BTNode{
	return &BinaryTree{}
}

//uから根への経路をたどる時のステップ数を返す
func (bt *BinaryTree) Depth(u *node) int {
	result := 0

	for u != nil{
		u = u.parent
		result++
	}

	return result
}

//ノードの数を返します。
func (bt *BinaryTree) Size(u *node) int{
	if u == nil {
		return 0
	}

	return  1 + Size(u.left) + Size(u.right)
}


