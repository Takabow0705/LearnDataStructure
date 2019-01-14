package chained_hash_table

type ChainedHashTable struct{
	n int
	array []string
}

func NewHashTable() *ChainHashTable{
	return &ChainHasTable{
		n : 0,
		array : make([]string,n + 1),
	}
}

