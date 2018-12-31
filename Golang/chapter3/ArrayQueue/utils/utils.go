package utils

type T int

func Compare(a,b int) int{
	return int(b-a)
}

func Max(a,b int) int{
	if Compare(a,b) < 0{
		return b
	}else{
		return a
	}
}