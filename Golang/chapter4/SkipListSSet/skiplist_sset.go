package skiplist_sset

import(
	"math/rand"
)

type node struct{
	x string
	height int
	node []*next
}

func pickHeight() int{
	r := rand.Int()
	height := 0

	for r&1 != 0{
		height++
		r >>= 1
	}
	return height
}