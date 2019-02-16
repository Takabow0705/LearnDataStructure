package main

import(
	"./binary_tree"
	"fmt"
)

//binary_tree.goは二分木の概要なので、メソッドのテストはほとんどできない
//また、データの追加等にも対応していない。
//これらはBinarySearchTreeで詳しく行う。
func main(){
	bt := binary_tree.NewBT()
	fmt.Println(bt.Size())
}