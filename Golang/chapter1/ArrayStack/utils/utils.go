package utils

type T int

func Compare(a,b T) int{
	return int(b-a)
}

func Max(a,b T) int{
	if Compare < 0{
		return b
	}else{
		return a
	}
}