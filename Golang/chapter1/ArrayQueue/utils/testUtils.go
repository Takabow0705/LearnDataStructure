package utils

import(
	"fmt"
	"testing"
)

func main(){
	test1 := Compare(0,1)
	test2 := Compae(-129,324)
	test3 := Compae(3,3)

	fmt.Println("Expect: 1 Answer: %v, Expect: 324 Answer: %v, Expect: 3 Answer: %v"
				,test1,test2,test3)
	
}