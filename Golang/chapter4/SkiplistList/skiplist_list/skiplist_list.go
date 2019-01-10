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