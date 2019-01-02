package dual_array_deque

import(
	"../../../utils"
)

type DualArrayDeque struct{
	n int
	front,back ArrayStack
}

func NewDualArrayDeque() DualArrayDeque {
	return DualArrayDeque{
		front: NewArrayStack(),
		back:  NewArrayStack(),
	}
}
//ふたつのArrayStackのサイズの合計を
//DualArrayDequeのサイズとする
func (da *DualArrayDeque) Size() int{
	return da.front.Size() + da.back.Size()
}

//frontとbackのArrayStackで
//なるべく要素数が等しくなるように要素を保持している。
//ゆえに引数のインデックス次第で取得対象のArrayStackが変わる。
func (da *DualArrayDeque) Get(i int) utils.T{
	if i < da.front.Size(){
		return da.front.Get(da.front.Size() - i - 1)
	}else{
		return da.back.Get(i - da.front.Size())
	}
}

//frontとbackのArrayStackで
//なるべく要素数が等しくなるように要素を保持している。
//ゆえに引数のインデックス次第で挿入対象のArrayStackが変わる。
func (da *DualArrayDeque) Set(i int,x utils.T) utils.T{
	if i < da.front.Size(){
		return da.front.Set(da.front.Size() - i - 1,x)
	}else{
		return da.back.Set(i - da.front.Size(),x)
	}
}

//frontとbackのどちらかに要素を追加する。
//追加後にfrontとbackで要素数がなるべく等しくなるように
//balance()を呼び出す。
//ArrayQueue同様に追加後はnを1増やす。
func (da *DualArrayDeque) Add(i int, x utils.T){
	if i < da.front.Size(){
		da.front.Add(da.front.Size() - i ,x)
	}else{
		da.back.Add(i - da.front.Size(),x)
	}

	da.balance()
	da.n++
}

//frontとbackのどちらかの要素を削除する。
//削除後にfrontとbackで要素数がなるべく等しくなるように
//balance()を呼び出す。
//ArrayQueue同様に追加後はnを1減らす。
func (da *DualArrayDeque) Remove(i int) utils.T{
	var result utils.T
	if i < da.front.Size(){
		result = da.front.Remove(da.front.Size() - i - 1)
	}else{
		result = da.back.Remove(i - da.front.Size())
	}
	da.balance()
	da.n--

	return result
}

//frontとbackの要素数の差が三倍以上となることが無いように
//frontとbackを調整する。
func (da *DualArrayDeque) balance(){
	front_size := da.front.Size()
	back_size := da.back.Size()

	if 3 * front_size < back_size || 3 * back_size < front_size{
		n := front_size + back_size
		half_n := n/2

		front_new := NewArrayStack()
		for i := half_n - 1; i >= 0; i-- {
			front_new.Push(da.Get(i))
		}

		back_new := NewArrayStack()
		for i := 0; i < half_n; i++ {
			back_new.Push(da.Get(half_n + i))
		}

		da.front = front_new
		da.back = back_new
	}
}
