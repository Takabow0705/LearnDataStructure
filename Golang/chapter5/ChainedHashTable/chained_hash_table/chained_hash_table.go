package chained_hash_table

import(
	"math/rand"
	"calc_utils"
)
type ChainedHashTable struct{
	n int
	array []ArrayStack
}

func NewHashTable() *ChainHashTable{
	return &ChainHashTable{
		n : 0,
		array : make([]ArrayStack,1<<1),
	}
}

func hash(x string) uint64{
	return calc_utils.MultipleHash(x,uint64(8))
}
func (self *ChainHashTable) Add(x string){

}