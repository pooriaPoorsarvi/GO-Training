package main

import "fmt"

func main(){

	var a [3]int
	lengthA := len(a)

	for i:=0; i < lengthA ; i ++ {
		fmt.Println(i, " : ", a[i])
	}

	for i, iv := range a {
		fmt.Println(i, " : ", iv)
	}

	fmt.Printf("Type of a is : %T and its value is %v \n", a, a)
	s := a[:]
	fmt.Printf("Type of s is : %T and its value is %v \n", s, s)

}
/*
output :
0  :  0
1  :  0
2  :  0
0  :  0
1  :  0
2  :  0
Type of a is : [3]int and its value is [0 0 0]
Type of s is : []int and its value is [0 0 0]
*/
