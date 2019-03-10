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

type ScapegoatTree struct {
	r *node
	n int
	q int //counter
}

func NewScapegoatTree() *ScapegoatTree {
	return &ScapegoatTree{}
}

//引数のnodeを根とした場合のサイズを返します。
func size_subtree(n *node) int {
	if n == nil {
		return 0
	}
	return size_subtree(n.left) + size_subtree(n.right) + 1
}
//二分木の要素数を返します。
func (st *ScapegoatTree) Size() int{
	return st.n
}
//二分探索木の探索を行い、探索結果を返します。
//存在するならデータを、しないならnilを返します。
func(st *ScapegoatTree) FindEQ(x int) interface{}{
	w := st.r

	for w != nil{
		comp := utils.IntCompare(x,w.x)

		if comp < 0{
			w = w.left
		}
		if comp > 0{
			w = w.right
		}

		if comp == 0{
			return w.x
		}
	}
	return nil
}

//探索を行い、探索の成否とデータを返す。
//true 	=> 探索は成功、もう一方の戻り値は探索結果
//false => 探索は失敗、もう一方の戻り値は最後に到達したnodeのデータ
func(st *ScapegoatTree) Find(x int) (bool,interface{}){
	w := st.r
	var z *node

	for w != nil{
		comp := utils.IntCompare(x,w.x)

		if comp < 0{
			z = w
			w = w.left
		}

		if comp > 0{
			w = w.left
		}

		if comp == 0{
			return true,w.x
		}

	}

	if z == nil {
		return false,nil
	}
	return false,w.x
}

//xを二分探索木に保存します。
func (st *ScapegoatTree) Add(x int) bool{
	p := st.findLast(x)
	u := newNode()
	u.x = x
	result := st.addChild(p,u)

	//初期状態はst.rもnilだからこの条件が必要
	if p == nil{
		st.r = u
		st.n++
		return true
	}

	if result {
		st.n++
		return result
	}
	return result
}

//探索の結果の最後のノードを返す。
func (st ScapegoatTree) findLast(x int) *node{
	w := st.r
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

//新しい葉を保存するメソッド
//保存が成功すれば true 失敗すれば false
func(st ScapegoatTree) addChild(p *node,u *node) bool{

	if p == nil{
		st.r = u
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
	st.n++
	return true
}

//nodeの削除後にツリーの連続性を確保します。
func(st *ScapegoatTree) splice(u *node){
	var s,p *node

	if u.left != nil{
		s = u.left
	}else{
		s = u.right
	}
	
	if u == st.r{
		st.r = s
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
	st.n--
}

func(st *ScapegoatTree) Remove(u *node){
	if u.left == nil || u.right == nil{
		st.splice(u)
		return
	}else{
		w := u.right

		for w.left != nil{
			w = w.left
			u.x = w.x
			st.splice(w)
		}
	}
}

//node u以下の葉の高さを再構築する。
//具体的には任意の葉の高さの差が高々1となるようにする。
func (st *ScapegoatTree) rebuild(u *node){
	ns := st.Size()
	p := u.parent
	a := make(*node,ns)

	st.packIntoArray(u,a,0)

	if p == nil {
		st.r = st.buildBalance(a,0,ns)
		st.r.parent = nil
	}else if p.right == u{
		p.right = buildBalance(a,0,ns)
		p.right.parent = p
	}else{
		p.left = buildBalance(a,0,ns)
		p.left.parent = p
	}
}

//uより下の葉を配列に格納し、nilとなったらその添字を返す。
func (st *ScapegoatTree) packIntoArray(u *node,a []*node,i int) int{
	if u == nil{
		return i
	}
	i = packIntoArray(u *node,a []*node,i int)
	i++
	a[i] = u
	return packIntoArray(u.right,a,i)
}
func (st *ScapegoatTree) buildBalance(a []*node, i int ,ns int) *node{
	if ns == 0{
		return nil
	}

	m := ns / 2
	a[i+m].left = st.buildBalanced(a, i, m)

	if a[i+m].left != nil {
		a[i+m].left.parent = a[i+m]
	}

	a[i+m].right = st.buildBalanced(a, i+m+1, ns-m-1)

	if a[i+m].right != nil {
		a[i+m].right.parent = a[i+m]
	}
	return a[i+m]
}

//スケープゴートツリーの木の高さを計算するための関数
//log_(3/2)を自然対数関数で表現した
func log32(n int) int {
	return int(math.Log(float64(n)) / math.Log(3.0/2))
}

//ツリーに新しい要素を追加する。
//その際、以下の条件を満たすかチェックする。
// st.q / 2 <= st.Size() <= st.q
func (st *ScapegoatTree) Add(x utils.V) bool {
	u := &node{x: x}
	d := st.addWithDepth(u)

	// st.q / 2 <= st.Size() <= st.q
	if d > log32(st.q) {
		w := u.parent
		a := size_subtree(w)
		b := size_subtree(w.parent)

		// (size(w.child) / size(w)) > 2/3
		for 3*a <= 2*b {
			w = w.parent
			a = size_subtree(w)
			b = size_subtree(w.parent)
		}
		st.rebuild(w.parent)
	} else if d < 0 {
		return false
	}
	return true
}