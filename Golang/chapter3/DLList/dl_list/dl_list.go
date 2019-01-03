package dl_list


type node struct{
	x string
	prev,next *node
}

//内部でnode構造体を生成するメソッド
func createNode(x string) *node{
	return &node{
		x : x,
	}
}

type DLList struct{
	n int
	dummy *node
}

func CreateDLList() *DLList{
	dummy := new(node)
	dummy.next = dummy
	dummy.prev = dummy

	return &DLList{
		dummy : dummy,
		n : 0,
	}
}

//与えられたindexのnodeを返す。
//indexの値がデータの要素数の中央の値からみて
//どちらに近いかで探索ルートを変更する。
func (dl *DLList) get_node(i int) *node{
	var p *node

	if i < dl.n/2 {
		p = dl.dummy.next
		for j := 0; j < i;j++{
			p = p.next
		}
	}else{
		p = dl.dummy.prev
		for j := 0; j < i;j++{
			p = p.prev
		}
	}
	return p
}

//指定されたindexに対応するnode構造体の要素を返す
func (dl *DLList) Get(i int) string{
	return dl.get_node(i).x
}

//指定されたindexに対応するnode構造体の要素に新しい要素を上書きする。
//その後、古い要素を返却する
func (dl *DLList) Set(i int, x string) string{
	u := dl.get_node(i)
	y := u.x
	u.x = x

	return y
}

//node wの直前にnode uを追加する。
//dummy のおかげでprevやnextの対象nodeが存在しなくても
//ロジックが複雑にはならない
func (dl *DLList) add_before(w *node,x string) *node{
	u := new(node)
	u.x = x
	u.prev = w.prev
	u.next = w

	u.prev.next = u
	u.next.prev = u

	dl.n++
	return u
}

func (dl *DLList) add_i(i int,x string){
	dl.add_before(dl.get_node(i),x)
}
//要素を先頭に追加
func (dl *DLList) Add(x string){
	dl.add_i(dl.n,x)
}

//p61上段の見本コードを実装
//Remove()の補助メソッド
func (dl *DLList) remove_i(i int) string{
	 p := dl.get_node(i)

	 u := p.x
	 p.prev.next = p.next
	 p.next.prev = p.prev
	 dl.n--
	 return u
}

//指定されたindexの要素を削除する
func (dl *DLList) Remove(i int) string{
	return dl.remove_i(i)
}

//最後の要素を削除する
func (dl *DLList) Pop() string{
	return dl.remove_i(dl.n - 1)
}

