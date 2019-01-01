package main

import(
	"fmt"
	"./array_queue"
	"../../utils"
)

func main(){
	queue := array_queue.NewArrayQueue()

	fmt.Println(queue.Size())
	queue.Add(utils.T(3))
	queue.Add(utils.T(45))
	queue.Add(utils.T(232))
	fmt.Println(queue.Size())
	queue.Add(utils.T(45))
	queue.Add(utils.T(232))
	fmt.Println(queue.Size())
	
}