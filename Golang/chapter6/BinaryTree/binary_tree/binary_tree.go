package binary_tree

type node struct {
	left   *node
	right  *node
	parent *node
}

func newNode() *node {
	retrun & node{}
}

type BinaryTree struct {
	r *node
}

func NewBTNOode() *BTNode {
	return &BinaryTree{}
}

//uから根への経路をたどる時のステップ数を返す
func (bt *BinaryTree) Depth(u *node) int {
	result := 0

	for u != nil {
		u = u.parent
		result++
	}

	return result
}

//ノードの数を返す、Size() メソッドの補助関数です
func (bt *BinaryTree) count_size_helper(u *node) int {
	if u == nil {
		return 0
	}

	return 1 + Size(u.left) + Size(u.right)
}

//ノードの数を返します。
func (bt *BinaryTree) Size() int {
	return count_size_helper(bt.r)
}

//ふたつの部分木の高さの最大値を返す、Height()メソッドの補助関数です
//nodeがはじめからnilの時は-1が返る。
func (bt *BinaryTree) measure_height_helper(u *node) int {
	if u == nil {
		return -1
	}

	return 1 + max(measure_height_helper(u.left), measure_height_helper(u.right))
}

//6.1.2の例示
//スタックオーバーフローを考慮していない
func (bt *BinaryTree) danger_traverse(u *node) {
	if u == nil {
		return
	}

	danger_traverse(u.left)
	danger_traverse(u.right)
}

//
func (bt *BinaryTree) Safety_traverse() {
	u := bt.r
	prev := nil
	var next *node

	for u != nil {
		if prev == u.parent && u.left != nil {
			next = u.left
		}

		if prev == u.parent && u.right != nil {
			next = u.right
		}

		if prev == u.parent && u.left == nil && u.right == nil {
			next = u.parent
		}

		if prev == u.left && u.right != nil {
			next = u.right
		}
		if prev == u.left && u.right == nil {
			next = u.parent
		}
		if prev == u.right {
			next = u.parent
		}

		prev = u
		u = next
	}
}
