package main

import(
	"fmt"
	"./ArrayStack"
	"./ArrayQueue"
)

func main(){
	stack := array_stack.NewArrayStack()
	queue := array_queue.NewArrayQueue()

	fmt.Println(stack.Size())
	stack.Add(0,3)
	stack.Add(1,45)
	stack.Add(2,22)
	fmt.Println(stack.Size())

	fmt.Println(queue.Size())
	queue.Add(0,3)
	queue.Add(1,45)
	queue.Add(2,22)
	fmt.Println(queue.Size())	
	
}