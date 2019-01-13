package skiplist_list

import(
	"math/rand"
	"strconv"
	"../../../utils"
)

type node struct{
	x int
	lengths int
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

func newNode(x int,h int,len int) *node{
	return &node{
		x : x,
		lengths: make([]int, h+1),
		height : h,
		nexts : make([]*node,h + 1),
	}
}

type SkipListList struct{
	sentinel *node
	n int
	height int
	stack     [strconv.IntSize]*node
}

func NewSkiplistList() *SkiplistList {
	sl := SkiplistList{
		sentinel: newNode(0, strconv.IntSize),
	}
	sl.stack[0] = ss.sentinel
	return &sl
}

func (sl *SkipListList) findPred(int i) *node{
	u := sl.sentinel
	r := sl.height
	j := -1

	for r>=0 {
		for u.nexts[r] != nil && j + u.lengths[r] < i{
			j += u.lengths[r]
			u = u.nexts[r]
		}
		r--
	}
	return u
}

//i番目のnodeを見つけてその要素を返す
func (sl *SkipListList) Get(i int)int{
	return sl.findPred(i).nexts[0].x
}

//i番目の要素に指定された要素に入れ替える
// i : 要素のインデックス
// x : 挿入する要素
func(sl *SkipListList) Set(i int,x int) int{
	u := sl.findPred(i).nexts[0]
	y := u.x
	u.x = x

	return y
}