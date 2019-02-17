package chained_hash_table

import(
	"math/rand"
	"calc_utils"
	"calc_utils/hash_type_enum"
)
type ChainedHashTable struct{
	n int
	array []ArrayStack
	hashtype HashType
}

func NewHashTable(hashtype HashType) *ChainHashTable{
	return &ChainHashTable{
		n : 0,
		array : make([]ArrayStack,1<<1),
		hashtype : hashtype
	}
}

func(cht *ChainHashTable) hash(x string) uint64{
	hash_func := calc_utils.HashTypeStrategy(cht.hashtype)
	return hash_func(x,uint64(8))
}

func (self *ChainHashTable) Add(x string){

}