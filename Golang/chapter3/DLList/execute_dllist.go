package main

import(
	"fmt"
	"./dl_list"
)

func main(){
	list := dl_list.CreateDLList()

	fmt.Println(list.Size())
	list.Add("hoge")
	list.Add("huga")
	list.Add("foo")
	fmt.Println(list.Size())
	fmt.Println(list.Pop())
	fmt.Println(list.Size())
	list.Add("foo")
	fmt.Println(list.Size())
}