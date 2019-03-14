package red_black_tree

type Colors uint

const(
	RED Colors = iota
	BLACK
)

type node struct{
	x int
	left *node
	right *node
	parent *node
	color Colors
}

func newNode(x int,color Colors) *node{
	return &node{
		x : x
		color : color
	}
}

type RedBlackTree struct{
	r *node
	n int
}

func NewRedBlackTree() *RedBlackTree{
	return &RedBlackTree{}
}

//子を黒に、uを赤にする
//uは赤の子を持つ黒ノードである
func (rbt *RedBlackTree) pushBlack(u *node){
	u.color = RED
	u.left.color = BLACK
	u.right.color = BLACK 
}

//子を赤に、uを黒にする
//uは黒の子を持つ赤ノードである。
func(rbt *RedBlackTree) pullBlack(u *node){
	u.color = BLACK
	u.left.color = RED
	u.right.color = RED
}

//uが左傾的で無いときに、左傾性を維持するために実行する。
//uが左傾的とは以下の条件を満たす時である。
//u.left.color = BLACK && u.right.color = RED
func(rbt *RedBlackTree) flipLeft(u *node){
	swapcolors(u,u.right)
	rbt.rotateRight()
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
	if w.parent != nil && w.parent.right == u{
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
		t.r.parent = nil
	}
}
