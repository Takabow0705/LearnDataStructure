package main

import(
	"./binary_search_tree"
	"fmt"
)

func main(){
	bst := binary_search_tree.NewBST()
	fmt.Println(bst.Add(5))
	fmt.Println(bst.Add(21))
	fmt.Println(bst.Add(2))
	fmt.Println(bst.Add(1))
	fmt.Println(bst.Add(3))
	fmt.Println(bst.Add(5))
	fmt.Println(bst.Size())
	fmt.Println(bst.Find(0))
	fmt.Println(bst.FindEQ(21))
	fmt.Println(&bst)
}