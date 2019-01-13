package utils

import (
	"reflect"
)
type T int

func Compare(a,b T) int{
	return int(b-a)
}

func IntCompare(a,b int)int{
	return b - a
}

func IsNil(x interface{}) bool{
	return x == nil && reflect.ValueOf(x).IsNil() 
}
func Max(a,b int) int{
	if a - b < 0{
		return b
	}else{
		return a
	}
}