package main

import "fmt"

func main(){
	s1 := []int{
		1,2,3,4,
	}
	s1[2] = 10
	// The following line will cause a problem no matter what you put instead of X
	//s1[2:4] = X

	fmt.Println(s1)
	fmt.Println(s1[2:4])

}

/*
output :
[1 2 10 4]
[10 4]
*/
