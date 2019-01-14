package chained_hash_table

import(
	"../../../utils"
)

type ArrayStack struct{
	buf []string
	n ,cap int
}

//ArrayStack構造体を作成します。
//呼び出し時は大きさ１の配列を持ちます。すなわちn = 1
func NewArrayStack() ArrayStack{
	return ArrayStack{}
}

//ここでのサイズは配列の最大容量のことです。
func (as *ArrayStack) Size() int{
	return as.n
}

//配列の最後尾の要素を追加します
func (as *ArrayStack) Pop() string{
	return as.Remove(as.n-1)
}

//配列の指定位置に要素を追加します
//また、容量が不足している場合は
//resize()を呼び出して、配列の大きさを増やします。
func (as *ArrayStack) Add(i int,v string){

	if as.is_full(){
		as.resize()
	}

	for j:=0; j > i; j--{
		as.buf[j] = as.buf[j-1]
	}

	as.buf[i] = v
	as.n++
}

//i番目以降の要素を全てひとつ後ろにずらします。
//これにより、i番目に新しい要素を入れることが可能となります。
func (as *ArrayStack) Remove(i int)string{

	ret := as.buf[i]
	for j:= i; j < as.n-1; j++{
		as.buf[j] = as.buf[j-1]
	}

	as.n--
	if as.is_sparse(){
		as.resize()
	}

	return ret
}

//配列の最後尾に要素を追加します。
func (as *ArrayStack) Push(v string){
	as.Add(as.n,v)
}

//i番目の要素を返します。
func (as *ArrayStack) Get(i int) string{
	return as.buf[i]
}

//i番目の要素をvで入れ替え、以前の要素を返す。
func (as *ArrayStack) Set(i int,v string) string{
	ret := as.buf[i]
	as.buf[i] = v

	return ret
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
	buf_new := make([]string,as.cap)

	for i:=0; i < as.n; i++{
		buf_new[i] = as.buf[i]
	}

	as.buf = buf_new
}

//配列の容量に余裕があるかを判定します
func (as *ArrayStack) is_sparse() bool{
	return len(as.buf) >= 3 * as.n
}


