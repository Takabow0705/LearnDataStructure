package array_deque

import (
	"./utils"
)

type ArrayDeque struct{
	n,j,cap int
	buf []utils.T
}

//ArrayDeque構造体を作成します。
//呼び出し時は大きさ１の配列を持ちます。すなわちn = 1
func NewArrayDeque() ArrayDeque{
	return ArrayDeque{}
}

//配列の指定位置に要素を追加します
//また、容量が不足している場合は
//resize()を呼び出して、配列の大きさを増やします。
func (as *ArrayDeque) Add(v utils.T){

	if as.is_full(){
		as.resize()
	}

	as.buf[(as.i + as.n)%as.cap] = v
	as.n++
}

//i番目以降の要素を全てひとつ後ろにずらします。
//これにより、i番目に新しい要素を入れることが可能となります。
func (as *ArrayDeque) Remove()utils.T{

	ret := as.buf[as.i]
	as.i = (as.i  + 1)%as.cap
	as.n--

	if as.is_sparse(){
		as.resize()
	}

	return ret
}

//ここでのサイズは配列の最大容量のことです。
func (as *ArrayDeque) Size() int{
	return as.n
}

//削除対象の要素を返します。
func (as *ArrayDeque) Get(i int) utils.T{
	return as.buf[(j + i)%as.cap]
}

//i番目の要素をvで入れ替え、以前の要素を返す。
func (as *ArrayDeque) Set(i int,v utils.T) utils.T{
	ret := as.buf[(j + i)%as.cap]

	as.buf[(j + i)%as.cap] = v

	return ret
}
//使用されている配列の要素数と配列の容量を比較します。
func (as *ArrayDeque) is_full() bool{
	return as.n == as.cap
}

//既存の配列の二倍の大きさをもつ配列を作成し,
//それまでのデータをコピーして
//新しい配列をレシーバの構造体の配列として再設定します。
func (as *ArrayDeque) resize(){
	cap_new := utils.Max(2 * as.n,1) 
	buf_new := make([]utils.T,cap_new)

	for i:=0; i < as.n; i++{
		buf_new[i] = as.buf[(i+as.i)%as.cap]
	}

	as.buf = buf_new
	as.cap = cap_new
	as.i = 0
}

//配列の容量に余裕があるかを判定します
func (as *ArrayDeque) is_sparse() bool{
	return len(as.buf) >= 3 * as.n
}