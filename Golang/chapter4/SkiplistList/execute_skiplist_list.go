package main

import(
	"fmt"
	"./skiplist_list"
)

func main(){
	list := skiplist_list.NewSkiplistList()
	fmt.Println(list.Size())
	list.Add(1,3)
	fmt.Println(list.Size())
	list.Remove(1)
	fmt.Println(list.Size())
}