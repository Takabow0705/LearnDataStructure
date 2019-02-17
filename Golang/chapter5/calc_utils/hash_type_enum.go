package calc_util

type HashType String

const(
	MultipleHash = iota
)

//列挙型のtoString()
func (self *HashType) String(h HashType) string{
	switch h {
	case MultipleHash :
		return "MultipleHash"
	}
}

