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

type BinarySearchTree struct {
	r *node
	n int
}

func NewBST() *BinarySearchTree {
	return &BinarySearchTree{}
}

//二分木の要素数を返します。
func (bst *BinarySearchTree) Size() int{
	return bst.n
}
//二分探索木の探索を行い、探索結果を返します。
//存在するならデータを、しないならnilを返します。
func(bst *BinarySearchTree) FindEQ(x int) interface{}{
	w := bst.r

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
func(bst *BinarySearchTree) Find(x int) (bool,interface{}){
	w := bst.r
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
func (bst *BinarySearchTree) Add(x int) bool{
	p := bst.findLast(x)
	u := newNode()
	u.x = x
	result := bst.addChild(p,u)

	//初期状態はbst.rもnilだからこの条件が必要
	if p == nil{
		bst.r = u
		bst.n++
		return true
	}

	if result {
		bst.n++
		return result
	}
	return result
}

//探索の結果の最後のノードを返す。
func (bst BinarySearchTree) findLast(x int) *node{
	w := bst.r
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
func(bst BinarySearchTree) addChild(p *node,u *node) bool{

	if p == nil{
		bst.r = u
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
	bst.n++
	return true
}

//nodeの削除後にツリーの連続性を確保します。
func(bst *BinarySearchTree) splice(u *node){
	var s,p *node

	if u.left != nil{
		s = u.left
	}else{
		s = u.right
	}
	
	if u == bst.r{
		bst.r = s
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
	bst.n--
}

func(bst *BinarySearchTree) Remove(u *node){
	if u.left == nil || u.right == nil{
		bst.splice(u)
		return
	}else{
		w := u.right

		for w.left != nil{
			w = w.left
			u.x = w.x
			bst.splice(w)
		}
	}
}