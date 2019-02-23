package main

import(
	"fmt"
	"./random_bst"
)

func main(){
	bst := random_bst.CreateTreap()
	fmt.Println(bst.Add(2))
	fmt.Println(bst.Add(4))
	fmt.Println(bst.Add(12))
	fmt.Println(bst.Add(3))
	fmt.Println(bst.Add(9))
	fmt.Println(bst.Add(5))
	fmt.Println(bst.Size())
	fmt.Println(bst.Remove(9))
	fmt.Println(bst.Remove(45))
	fmt.Println(bst.Size())
}