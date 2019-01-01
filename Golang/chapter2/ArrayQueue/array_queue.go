package array_queue

import(
	"./utils"
)

type ArrayQueue struct{
	buf []utils.T
	n ,cap ,i int
}

//ArrayQueue構造体を作成します。
//呼び出し時は大きさ１の配列を持ちます。すなわちn = 1
func NewArrayQueue() ArrayQueue{
	return ArrayQueue{}
}

//配列の指定位置に要素を追加します
//また、容量が不足している場合は
//resize()を呼び出して、配列の大きさを増やします。
func (aq *ArrayQueue) Add(v utils.T){

	if aq.is_full(){
		aq.resize()
	}

	aq.buf[(aq.i + aq.n)%aq.cap] = v
	aq.n++
}

//i番目以降の要素を全てひとつ後ろにずらします。
//これにより、i番目に新しい要素を入れることが可能となります。
func (aq *ArrayQueue) Remove()utils.T{

	ret := aq.buf[aq.i]
	aq.i = (aq.i  + 1)%aq.cap
	aq.n--

	if aq.is_sparse(){
		aq.resize()
	}

	return ret
}

//ここでのサイズは配列の最大容量のことです。
func (aq *ArrayQueue) Size() int{
	return aq.n
}

//削除対象の要素を返します。
func (aq *ArrayQueue) Get() utils.T{
	return aq.buf[i]
}
//使用されている配列の要素数と配列の容量を比較します。
func (aq *ArrayQueue) is_full() bool{
	return aq.n == aq.cap
}

//既存の配列の二倍の大きさをもつ配列を作成し,
//それまでのデータをコピーして
//新しい配列をレシーバの構造体の配列として再設定します。
func (aq *ArrayQueue) resize(){
	cap_new := utils.Max(2 * aq.n,1) 
	buf_new := make([]utils.T,cap_new)

	for i:=0; i < aq.n; i++{
		buf_new[i] = aq.buf[(i+aq.i)%aq.cap]
	}

	aq.buf = buf_new
	aq.cap = cap_new
	aq.i = 0
}

//配列の容量に余裕があるかを判定します
func (aq *ArrayQueue) is_sparse() bool{
	return len(aq.buf) >= 3 * aq.n
}


