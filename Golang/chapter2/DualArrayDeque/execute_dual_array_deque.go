package main

import (
	"../../utils"
	"./dual_array_deque"
	"fmt"
)

func main(){

	dad := dual_array_deque.NewDualArrayDeque()

	//ToDo not use resize() and balance()?
	fmt.Println(dad.Size())
	dad.Add(0,string(3))
	dad.Add(0,string(45))
	dad.Add(0,string(232))
	fmt.Println(dad.Size())
	dad.Add(0,string(45))
	dad.Add(0,string(232))
	fmt.Println(dad.Size())
}