package main

import(
	"fmt"
	"./array_deque"
	"../../utils"
)

func main(){
	deque := array_deque.NewArrayDeque()

	fmt.Println(deque.Size())
	deque.Add(utils.T(3))
	deque.Add(utils.T(45))
	deque.Add(utils.T(232))
	fmt.Println(deque.Size())
	deque.Add(utils.T(45))
	deque.Add(utils.T(232))
	fmt.Println(deque.Size())
	
}