package main

import (
	"fmt"
	"./skiplist_sset"
)

func main(){

	list := skiplist_sset.NewSkipListSSet() 
	fmt.Println(list.Size())
	list.Add(5)
	list.Add(3)
	list.Add(12)
	fmt.Println(list.Size())
	list.Remove(1)
	fmt.Println(list.Size())

}