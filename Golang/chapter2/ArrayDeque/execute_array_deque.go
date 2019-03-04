package main

import(
	"fmt"
	"./array_deque"
)

func main(){
	deque := array_deque.NewArrayDeque()

	fmt.Println(deque.Size())
	deque.Add(0,string(3))
	deque.Add(3,string(45))
	deque.Add(5,string(232))
	fmt.Println(deque.Size())
	deque.Add(3,string(45))
	deque.Add(2,string(232))
	fmt.Println(deque.Size())
	fmt.Println(deque.String())
	
}