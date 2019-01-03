package dl_list


type node struct{
	x string
	prev,next *node
}

//内部でnode構造体を生成するメソッド
func createNode(x) *node{
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
	u := get_node(i)
	y := u.x
	u.x = x

	return y
}



