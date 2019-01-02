package main

import (
	"../../utils"
	"./dual_array_deque"
	"fmt"
)

func main(){

	dad := dual_array_deque.NewDualArrayDeque()

	fmt.Println(dad.Size())
	dad.Add(0,utils.T(3))
	dad.Add(0,utils.T(45))
	dad.Add(0,utils.T(232))
	fmt.Println(dad.Size())
	dad.Add(0,utils.T(45))
	dad.Add(0,utils.T(232))
	fmt.Println(dad.Size())
}