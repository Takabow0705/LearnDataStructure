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

func (st *ScapegoatTree) rebuild(u *node){
	ns := st.Size()
	p := u.parent
	a := make(*node,ns)

	st.packIntoArray(u,a,0)

	if p == nil {
		r = st.buildBalance(a,0,ns)
		r.parent = nil
	}else if p.right == u{
		p.right = buildBalance(a,0,ns)
		p.right.parent = p
	}else{
		p.left = buildBalance(a,0,ns)
		p.left.parent = p
	}
}

func (st *ScapegoatTree) packIntoArray(u *node,a.[]*node,i int) int{
	if u == nil{
		return i
	}
	i = packIntoArray(u *node,a []*node,i int)
	i++
	a[i] = u
	return packIntoArray(u.right,a,i)
}