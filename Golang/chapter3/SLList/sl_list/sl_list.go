package sl_list

import(
	../../utils
)

type node struct{
	x string
	next *node
}

func newNode(x string) *node {
	return &node{
		x: x,
	}
}

type SLList struct{
	n int
	head *node
	tail *node
}

func newSLList() *SLList{
	return new(SLList)
}

//SLListの大きさを返します。
func (sl *SLList) Size(){
	return sl.n
}

//新しいノードを作成し、それに新しいheadを追加する
//追加後はnを1増やす。
//またもしn = 0ならば新しいnodeはtailにも設定する。
func (sl *SLList) Push(x string){

	node_new := newNode(x)
	node_new.next = sl.head
	sl.head = node_new

	if n == 0{
		tail = node_new
	}
	sl.n++
}

//要素を取り出した後削除
//要素がない場合はnilを返す。
//n = 1ならtail はnil
func(sl *SLList) Pop() string{
	if sl.n == 0{
		return nil
	}

	u := sl.head.x
	sl.head = sl.head.next
	
	if sl.n == 1 {
		sl.tail = nil
	}
	sl.n--

	return u
}

//popメソッドのラッパー
func (sl *SLList) Remove(){
	sl.Pop()
}

//新しいノードを作成し、それに新しいheadを追加する
//追加後はnを1増やす。
//またもしn = 0ならば新しいnodeはtailにも設定する。
func (sl *SLList) Add(x string){
	node_new := newNode(x)

	if n == 0{
		sl.head = node_new
	}else{
		sl.tail.next = node_new
	}

	sl.tail = node_new
	sl.n++
}