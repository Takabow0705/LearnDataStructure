package main

import(
	"fmt"
	"./array_deque"
	"../../utils"
)

func main(){
	deque := array_deque.NewArrayDeque()

	fmt.Println(deque.Size())
	deque.Add(0,utils.T(3))
	deque.Add(3,utils.T(45))
	deque.Add(5,utils.T(232))
	fmt.Println(deque.Size())
	deque.Add(3,utils.T(45))
	deque.Add(2,utils.T(232))
	fmt.Println(deque.Size())
	
}