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

func createTreap() *Treap{
	return &Treap{
		n : 0
	}
}

//二分探索木の左回転を実装する。
func (t *Treap) rotateLeft(u *node){
	w := u.right
	w.parent = u.parent

	//uの親ノードの子ノードをwに切り替える。
	//左右どちらにあるかで場合分け
	if w.parent != nil && w.parent.left == u{
		w.parent.left = w
	}
	if w.parent != nil && w.parent.right == u{{
		w.parent.right = w
	}
	
	u.right = w.left

	//ノードuに付け替えた、ノードwの左子要素の親をノードuに設定する。
	if u.right != nil{
		u.right.parent = u
	}
	u.parent = w
	w.left = u

	//uがrootノードの時の処理
	if u == t.r {
		t.r = w
		t.r.parent = nil
	}
}

//二分探索木の右回転を実装
//実装の詳細は左回転の場合と同じ
func (t *Treap) rotateRight(u *node){
	w := u.left
	w.parent = u.parent

	if w.parent != nil && w.parent.left == u{
		w.parent.left = w
	}
	if w.parent != nil && w.parent.right == u{
		w.parent.right = w
	}

	u.left = w.right

	if u.left != nil{
		u.left.parent = u
	}
	u.parent = w
	w.right = u

	if u == t.r{
		t.r = w
		r.parent = nil
	}
}
