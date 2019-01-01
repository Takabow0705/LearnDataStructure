package main

import(
	"fmt"
	"./array_queue"
	"../../utils"
)

func main(){
	stack := array_queue.NewArrayStack()

	fmt.Println(stack.Size())
	stack.Add(0,utils.T(3))
	stack.Add(1,utils.T(45))
	stack.Add(2,utils.T(232))
	fmt.Println(stack.Size())
	stack.Add(3,utils.T(45))
	stack.Add(5,utils.T(232))
	fmt.Println(stack.Size())
	
}