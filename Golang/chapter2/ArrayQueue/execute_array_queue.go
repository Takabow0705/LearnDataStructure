package main

import(
	"fmt"
	"./array_queue"
	"../../utils"
)

func main(){
	stack := array_queue.NewArrayQueue()

	fmt.Println(stack.Size())
	stack.Add(utils.T(3))
	stack.Add(utils.T(45))
	stack.Add(utils.T(232))
	fmt.Println(stack.Size())
	stack.Add(utils.T(45))
	stack.Add(utils.T(232))
	fmt.Println(stack.Size())
	
}