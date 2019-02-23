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

func createNode(x int) *node{
	return &node{
		x : x
	}
}
type struct Treap{
	r *node
	n int
}

func CreateTreap() *Treap{
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

//探索の結果の最後のノードを返す。
func (t *Treap) findLast(x int) *node{
	w := t.r
	var prev *node
	
	for w != nil{
		prev = w
		comp := utils.IntCompare(x,w.x)
		
		if comp < 0{
			w = w.left
		}

		if comp > 0{
			w = w.right
		}

		if comp == 0{
			return w
		}

	}
	return prev
}

//nodeの削除後にツリーの連続性を確保します。
func(t *Treap) splice(u *node){
	var s,p *node

	if u.left != nil{
		s = u.left
	}else{
		s = u.right
	}
	
	if u == t.r{
		t.r = s
	}else{
		p = u.parent

		if &p.left == &u{
			p.left = s
		}else{
			p.right = s
		}
	}

	if s != nil{
		s.parent = p
	}
	t.n--
}

//新しい葉を保存するメソッド
//保存が成功すれば true 失敗すれば false
func(t *Treap) addChild(p *node,u *node) bool{

	if p == nil{
		t.r = u
	}else{
		comp := utils.IntCompare(u.x,p.x)

		if comp < 0{
			p.left = u
		}else if comp > 0{
			p.right = u
		}else{
			return false
		}
		u.parent = p
	}
	t.n++
	return true
}
