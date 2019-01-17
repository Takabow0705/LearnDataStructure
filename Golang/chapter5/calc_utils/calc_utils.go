package calc_utils

import(
	"math/rand"
	"strconv"
)
type func(x string,d uint64) uint64 HashCalcStrategy

//乗算ハッシュ法を実装します。 
//d : 次数
//x : 要素
//具体的には以下の計算を実装する。
// hash(x,d) = ( (z * uint64(x)) mod 2^w ) / 2^(w-d)
//この場合 w = 64である。
func MultipleHash(x string,d uint64) uint64{
	return (rand.Uint64*Uint64(x)) >> (64 - d)
}

func stringTouint64(x string) uint64{
	bytes := []byte(x)
	result := binary.LittleEndian.Uint16(bytes)
	return uint64(result)
}



