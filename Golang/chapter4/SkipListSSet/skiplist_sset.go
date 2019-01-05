package skiplist_sset

import(
	"math/rand"
	"strconv"
)

type node struct{
	x string
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

func newNode(x string,h int) *node{
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
		sentinel : newNode("", strconv.IntSize), //ToDo 初期値は""で良いのか
	}
	s.stack[0] = s.sentinel
	return &s
}