package utils

import(
	"fmt"
)

func main(){
	test1 := Compare(0,1)
	test2 := Compare(-129,324)
	test3 := Compare(3,3)

	fmt.Printf("Expect: 1 Answer: %v, Expect: 324 Answer: %v, Expect: 3 Answer: %v",test1,test2,test3)
	
}