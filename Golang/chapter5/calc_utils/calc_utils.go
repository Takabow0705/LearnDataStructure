package calc_utils

import(
	"math/rand"
	"strconv"
	"calc_utils/hash_type_enum"
)
type func(x string,d uint64) uint64 HashFunc

func HashTypeStrategy(hashtype HashType) HashFunc{
	switch hashtype{
	case MultipleHash:
		//MultipleHashの無名関数
		//下記実装を参照
		return func(x string,d uint64){
			return (rand.Uint64*stringTouint64(x)) >> (64 - d)
		}
	}
}

//乗算ハッシュ法を実装します。 
//d : 次数
//x : 要素
//具体的には以下の計算を実装する。
// hash(x,d) = ( (z * uint64(x)) mod 2^w ) / 2^(w-d)
//この場合 w = 64である。
func MultipleHash(x string,d uint64) uint64{
	return (rand.Uint64*stringTouint64(x)) >> (64 - d)
}

func stringTouint64(x string) uint64{
	bytes := []byte(x)
	result := binary.LittleEndian.Uint16(bytes)
	return uint64(result)
}



