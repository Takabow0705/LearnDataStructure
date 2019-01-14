package chained_hash_table

type ChainedHashTable struct{
	n int
	array ArrayStack
}

func NewHashTable() *ChainHashTable{
	return &ChainHashTable{
		n : 0,
		array : make([]ArrayStack,1<<1),
	}
}

func (self *ChainHashTable) Add(x string){

}