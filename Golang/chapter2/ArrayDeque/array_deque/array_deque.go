package array_deque

import (
	"fmt"
	"../../../utils"
)

type ArrayDeque struct{
	n,i,cap int
	buf []string
}

//ArrayDeque構造体を作成します。
//呼び出し時は大きさ１の配列を持ちます。すなわちn = 1
func NewArrayDeque() ArrayDeque{
	return ArrayDeque{}
}

//配列の指定位置に要素を追加します
//また、容量が不足している場合は
//resize()を呼び出して、配列の大きさを増やします。
func (ad *ArrayDeque) Add(i int ,v string){

	if ad.is_full(){
		ad.resize()
	}

	if i < ad.n/2{
		ad.i = ad.select_index()
		for j := 0; j < i; j++ {
			ad.buf[(j + i)%ad.cap] = ad.buf[(j + i + 1)%ad.cap]
		}
	}else{
		for j := ad.n; j > i; j-- {
			ad.buf[(j + i)%ad.cap] = ad.buf[(j + i - 1)%ad.cap]
		}
	}
	ad.buf[(ad.i + i)%ad.cap] = v
	ad.n++
}

//i番目以降の要素を全てひとつ後ろにずらします。
//これにより、i番目に新しい要素を入れることが可能となります。
func (ad *ArrayDeque) Remove(i int)string{

	ret := ad.buf[(i + ad.i)%ad.cap]
	if i < ad.n/2{
		for k := i; k > 0; k--{
			ad.buf[(i + k)%ad.cap] = ad.buf[(i + k - 1)%ad.cap]
		}
		ad.i = (i + 1)%ad.cap
	}else{
		for k := i; k < ad.cap; k++{
			ad.buf[(i + k)%ad.cap] = ad.buf[(i + k - 1)%ad.cap]
		}
	}
	ad.n--

	if ad.is_sparse(){
		ad.resize()
	}

	return ret
}

//ここでのサイズは配列の最大容量のことです。
func (ad *ArrayDeque) Size() int{
	return ad.cap
}

//削除対象の要素を返します。
func (ad *ArrayDeque) Get(i int) string{
	return ad.buf[(ad.i + i)%ad.cap]
}

//i番目の要素をvで入れ替え、以前の要素を返す。
func (ad *ArrayDeque) Set(i int,v string) string{
	ret := ad.buf[(ad.i + i)%ad.cap]

	ad.buf[(ad.i + i)%ad.cap] = v

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
	buf_new := make([]string,cap_new)

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
//Add(i int,x string)の補助メソッド
func (ad *ArrayDeque) select_index() int{
	if ad.i == 0{
		return ad.cap - 1
	}else{
		return ad.i - 1
	}
}

//データ構造を文字列として出力します。
func (ad *ArrayDeque) String() string{
	result := ""
	size := ad.Size()

	for i:=0; i < size; i++{
		str := fmt.Sprintf("{%d:%s},\n",i,ad.Get(i))
		result += str
	}
	return result
}