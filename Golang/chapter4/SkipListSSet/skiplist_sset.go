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
