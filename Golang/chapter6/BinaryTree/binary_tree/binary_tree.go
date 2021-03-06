package binary_tree

import(
	"../../../utils"
)

type node struct {
	left   *node
	right  *node
	parent *node
}

func newNode() *node {
	return & node{}
}

type BinaryTree struct {
	r *node
}

func NewBT() *BinaryTree {
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

	return 1 + bt.count_size_helper(u.left) + bt.count_size_helper(u.right)
}

//ノードの数を返します。
func (bt *BinaryTree) Size() int {
	return bt.count_size_helper(bt.r)
}

//ふたつの部分木の高さの最大値を返す、Height()メソッドの補助関数です
//nodeがはじめからnilの時は-1が返る。
func (bt *BinaryTree) measure_height_helper(u *node) int {
	if u == nil {
		return -1
	}

	return 1 + utils.Max(bt.measure_height_helper(u.left), bt.measure_height_helper(u.right))
}

//6.1.2の例示
//スタックオーバーフローを考慮していない
/*
func (bt *BinaryTree) danger_traverse(u *node) {
	if u == nil {
		return
	}

	danger_traverse(u.left)
	danger_traverse(u.right)
}
*/

//直前のnodeによって次のnodeを決めることで
//再帰を使用せずに二分木を走査する。
//以下の規則によって行き先を決定する
//
//prev : 以前のnode
//next : 次の走査対象
//u	   : 現在のnode
//
//規則:
//prev == parent   => next = u.left
//prev == u.left   => next = u.right
//prev == u.right  => next = u.left
func (bt *BinaryTree) Safety_traverse() int {
	u := bt.r
	n := 0
	var next *node
	var prev *node	

	for u != nil {
		if prev == u.parent && u.left != nil {
			next = u.left
			n++
		}

		if prev == u.parent && u.right != nil {
			next = u.right
			n++
		}

		if prev == u.parent && u.left == nil && u.right == nil {
			next = u.parent
			n++
		}

		if prev == u.left && u.right != nil {
			next = u.right
			n++
		}
		if prev == u.left && u.right == nil {
			next = u.parent
			n++
		}
		if prev == u.right {
			next = u.parent
			n++
		}
		if prev == nil{
			next = u.left
		}

		prev = u
		u = next
	}
	return n
}

//走査の結果として訪問したnodeの数を返す
//デフォルトは0
func(bt *BinaryTree) SizeByTraverse() int{
	result := bt.Safety_traverse()
	return result
}

//幅優先探索による走査の実装
func(bt *BinaryTree) BF_Traverse() int{
	n := 0
	q := make([]*node,0)

	if bt.r != nil{
		q = append(q,bt.r)
	}

	for len(q) > 0{
		//キューを模倣する
		//先頭の要素を取り出し、その後先頭要素を除いた配列を再代入する
		u := q[0]
		q = q[1:]

		if u.left != nil { 
			q = append(q,u.left) 
			n++ 
		}

		if u.right != nil { 
			q = append(q,u.right) 
			n++ 
		}
	}

	return n
}
