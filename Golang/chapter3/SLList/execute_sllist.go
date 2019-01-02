package main

import(
	"./sl_list"
	"fmt"
)

func main(){
	list := sl_list.NewSLList();

	fmt.Println(list.Size())
	list.Add("hoge")
	list.Add("huga")
	list.Add("hoo")
	fmt.Println(list.Size())
	fmt.Println(list.Pop())
	fmt.Println(list.Size())
}