package skiplist_sset

import(
	"math/rand"
)

type node struct{
	x string
	height int
	node *[]next
}