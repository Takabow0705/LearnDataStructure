package array_stack

import(
	"./utils"
)

type ArrayStack struct{
	buffer []utils.T
	n ,cap int
}

func NewArrayStack(len int) *ArrayStack{
	return *ArrayStack
}

//ここでのサイズは配列の最大容量のことです。
func (as *ArrayStack) Size() int{
	return as.n
}
func (as *ArrayStack) Add(i int,v utils.V){

	if as.is_full(){
		as.resize()
	}


}
//
func (as *ArrayStack) Push(v utils.T){
	as.Add(as.n,v)
}

//使用されている配列の要素数と配列の容量を比較します。
func (as *ArrayStack) is_full() bool{
	return as.n == as.cap
}

//既存の配列の二倍の大きさをもつ配列を作成し,
//それまでのデータをコピーして
//新しい配列をレシーバの構造体の配列として再設定します。
func (as *ArrayStack) resize(){
	as.cap = utils.Max(2 * as.n,1) 
	buf_new := make([]utils.T,as.cap)

	for i:=0; i < as.n; i++{
		buf_new[i] = as.buf[i]
	}

	as.buf = buf_new
}

//配列の容量に余裕があるかを判定します
func (as *ArrayStack) is_sparse() bool{
	return len(as.buf) >= 3 * as.n
}


