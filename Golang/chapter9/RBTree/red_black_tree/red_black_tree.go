package red_black_tree

type Colors uint

const(
	RED Colors = iota
	BLACK
)

type node struct{
	x int
	left *node
	right *node
	parent *node
	color Colors
}

func newNode(x int,color Colors) *node{
	return &node{
		x : x
		color : color
	}
}

type RedBlackTree struct{
	r *node
	n int
}

func NewRedBlackTree() *RedBlackTree{
	return &RedBlackTree{}
}

//子を黒に、uを赤にする
//uは赤の子を持つ黒ノードである
func (rbt *RedBlackTree) pushBlack(u *node){
	u.color = RED
	u.left.color = BLACK
	u.right.color = BLACK 
}

//子を赤に、uを黒にする
//uは黒の子を持つ赤ノードである。
func(rbt *RedBlackTree) pullBlack(u *node){
	u.color = BLACK
	u.left.color = RED
	u.right.color = RED
}

