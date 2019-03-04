package main

import(
	"fmt"
	"./array_queue"
)

func main(){
	queue := array_queue.NewArrayQueue()

	fmt.Println(queue.Size())
	queue.Add("3")
	queue.Add("45")
	queue.Add("232")
	fmt.Println(queue.Size())
	queue.Add("45")
	queue.Add("332")
	fmt.Println(queue.Size())
	fmt.Println(queue.String())
	
}