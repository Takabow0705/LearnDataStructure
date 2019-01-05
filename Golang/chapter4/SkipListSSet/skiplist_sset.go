package skiplist_sset

import(
	"math/rand"
)

type node struct{
	x string
	height int
	next []*node
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
		next : make([]*node,h + 1),
	}
}