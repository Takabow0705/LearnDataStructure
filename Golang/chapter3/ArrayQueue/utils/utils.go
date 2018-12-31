package utils

type T int

func Compare(a,b T) int{
	return int(b-a)
}

func Max(a,b T) int{
	if Compare(a,b) < 0{
		return int(b)
	}else{
		return int(a)
	}
}