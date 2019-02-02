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

//ノードの数を返す、Size() メソッドの補助関数です
func (bt *BinaryTree) count_size_helper(u *node) int{
	if u == nil {
		return 0
	}

	return  1 + Size(u.left) + Size(u.right)
}

//ノードの数を返します。
func (bt *BinaryTree) Size() int{
	retrun count_size_helper(bt.r)
}

//ふたつの部分木の高さの最大値を返す、Height()メソッドの補助関数です
//nodeがはじめからnilの時は-1が返る。
func (bt *BinaryTree) measure_height_helper(u *node) int{
	if u == nil{
		return -1
	}

	return 1 + max(measure_height_helper(u.left),measure_height_helper(u.right))
}

