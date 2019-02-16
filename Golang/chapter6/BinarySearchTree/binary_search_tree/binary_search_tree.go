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

func NewBT() *BinarySearchTree {
	return &BinarySearchTree{}
}

//二分探索木の探索を行い、探索結果を返します。
//存在するならデータを、しないならnilを返します。
func(bst *BinarySearchTree) FindEQ(x int) interface{}{
	w := bst.r

	for w != nil{
		comp := utils.compare(x,w.x)

		if comp < 0{
			w = w.left
		}
		if comp > 0{
			w = w.right
		}

		if comp == 0{
			return w.x
		}
		return nil
	}
}

//探索を行い、探索の成否とデータを返す。
//true 	=> 探索は成功、もう一方の戻り値は探索結果
//false => 探索は失敗、もう一方の戻り値は最後に到達したnodeのデータ
func(bst *BinarySearchTree) Find(x int) (bool,interface{}){
	w := bst.r
	var z *node

	for w != nil{
		comp := utils.Compare(x,z.x)

		if comp < 0{
			z = w
			w = w.left
		}

		if comp > 0{
			w = w.left
		}

		if comp == 0{
			return (true,w.x)
		}

	}

	if z == nil {
		return (false,nil)
	}
	return (false,w.x)
}

