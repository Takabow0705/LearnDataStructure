package array_stack

import(
	"./utils"
)

type ArrayStack struct{
	buffer []utils.T
	length ,cap int
}

func NewArrayStack(len int) *ArrayStack{
	return *ArrayStack
}

//ここでのサイズは配列の最大容量のことです。
func (as *ArrayStack) Size() int{
	return as.n
}

//使用されている配列の要素数と配列の容量を比較します。
//
func (as *ArrayStack) is_full() bool{
	return as.n == as.cap
}


