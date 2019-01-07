package skiplist_sset

import(
	"math/rand"
	"strconv"
	"../../utils"
)

type node struct{
	x int
	height int
	nexts []*node
}

//新しいnodeの高さを乱数で設定します。
//rの二進数表記で1の数をカウントする。
func pickHeight() int{
	r := rand.Int()
	height := 0

	for r&1 != 0{
		height++
		r >>= 1
	}
	return height
}

func newNode(x int,h int) *node{
	return &node{
		x : x,
		height : h,
		nexts : make([]*node,h + 1),
	}
}

//strconv.IntSize() は　bit or uintのサイズを返す
type SkipListSSet struct{
	sentinel *node
	n int
	height int
	stack     [strconv.IntSize]*node
}

func NewSkipListSSet() *SkipListSSet{
	s := SkipListSSet{
		sentinel : newNode(0, strconv.IntSize), //ToDo 初期値は""で良いのか
	}
	s.stack[0] = s.sentinel
	return &s
}

//探索経路に沿ってy>=xとなるyを含むnodeを返します。
//
func (ss *SkipListSSet) findPredNode(x int) *node{
	u := ss.sentinel
	r = ss.height

	for r >= 0{
		for u.nexts[r] != nil && Compare(u.nexts[r].x,x){
			u = u.nexts[r]
		}
		r--
	}
	return u
}

//探索経路に沿ってy>=xとなるyを含むyと探索結果を返します。
//ok : 探索結果　true => 成功　:  false => 失敗
//result : 探索結果の値
func (ss *SkipListSSet) Find(x int) (ok bool ,result int){
	u := ss.findPredNode(x)
	if u.nexts[0] == nil{
		result = u.nexts[0].x
		ok = true
		return ok,result
	}else{
		result = 0
		ok = false
		return  ok,result
	}
	
}

//
func(ss * SkipListSSet) Add(x int) bool{
	u := ss.sentinel
	r := ss.height
	comp := 0

	for r >= 0{
		for u.nexts[r] != nil && utils.Compare(u.nexts[r].x, x) < 0 {
			u = u.nexts[r]
		}
		if u.nexts[r] != nil && utils.Compare(u.nexts[r].x, x) == 0 {
			return false
		}
		ss.stack[r] = u
		r--
	}

	w := ss.newNode(x,pickHeight())
	for ss.height < w.height {
		ss.height++
		ss.stack[ss.height] = ss.sentinel
	}
	for i := 0; i <= w.height; i++ {
		w.nexts[i] = ss.stack[i].nexts[i]
		ss.stack[i].nexts[i] = w
	}
	ss.n++
	return true
}

//
func (ss *SkipListSSet) remove(x int) bool{
	ok  = false
	u := ss.sentinel
	r := height
	comp := 0

	for r >= 0{
		for u.nexts[r] != nil && utils.Compare(u.nexts[r].x,x) < 0{
			u = u.nexts[r]
		}

		if u.nexts[r] != nil && utils.Compare(u.nexts[r].x,x) == 0{
			ok = true
			del = u.nexts[r]

			u.nexts[r] = u.nexts[r].nexts[r]
			if u == ss.sentinel && u.nexts[r] != nil{
				height--
			}
		}
		r--
	}
	return ok
}
