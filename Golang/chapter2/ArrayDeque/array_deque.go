package array_deque

import (
	"github.com/Takabow0705/LearnDataStructure/tree/master/Golang/utils"
)

type ArrayDeque struct{
	n,i,cap int
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
func (ad *ArrayDeque) Add(i int ,v utils.T){

	if ad.is_full(){
		ad.resize()
	}

	if i < ad.n/2{
		ad.i = select_index()
		for j := 0; j < i; j++ {
			ad.buf[(j + i)%as.cap] = ad.buf[(j + i + 1)%as.cap]
		}
	}else{
		for j := ad.n; j > i; j--{
			ad.buf[(j + i)%as.cap] = ad.buf[(j + i - 1)%as.cap]
		}
	}
	ad.buf[(ad.i + i)%ad.cap] = v
	ad.n++
}

//i番目以降の要素を全てひとつ後ろにずらします。
//これにより、i番目に新しい要素を入れることが可能となります。
func (ad *ArrayDeque) Remove()utils.T{

	ret := ad.buf[ad.i]
	ad.i = (ad.i  + 1)%ad.cap
	ad.n--

	if ad.is_sparse(){
		ad.resize()
	}

	return ret
}

//ここでのサイズは配列の最大容量のことです。
func (ad *ArrayDeque) Size() int{
	return ad.n
}

//削除対象の要素を返します。
func (ad *ArrayDeque) Get(i int) utils.T{
	return ad.buf[(j + i)%ad.cap]
}

//i番目の要素をvで入れ替え、以前の要素を返す。
func (ad *ArrayDeque) Set(i int,v utils.T) utils.T{
	ret := ad.buf[(j + i)%ad.cap]

	ad.buf[(j + i)%ad.cap] = v

	return ret
}
//使用されている配列の要素数と配列の容量を比較します。
func (ad *ArrayDeque) is_full() bool{
	return ad.n == ad.cap
}

//既存の配列の二倍の大きさをもつ配列を作成し,
//それまでのデータをコピーして
//新しい配列をレシーバの構造体の配列として再設定します。
func (ad *ArrayDeque) resize(){
	cap_new := utils.Max(2 * ad.n,1) 
	buf_new := make([]utils.T,cap_new)

	for i:=0; i < ad.n; i++{
		buf_new[i] = ad.buf[(i+ad.i)%ad.cap]
	}

	ad.buf = buf_new
	ad.cap = cap_new
	ad.i = 0
}

//配列の容量に余裕があるかを判定します
func (ad *ArrayDeque) is_sparse() bool{
	return len(ad.buf) >= 3 * ad.n
}

//削除対象インデックスの値に応じてインデックスを再設定する
//Add(i int,x utils.T)の補助メソッド
func (ad *ArrayDeque) select_index() int{
	if i == 0{
		return ad.cap - 1
	}else{
		return i - 1
	}
}