package main

import(
	"fmt"
	"./skiplist_list"
)

func main(){
	list := skiplist_list.NewSkiplistList()
	fmt.Println(list.Size())
}