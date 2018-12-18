package main

import "fmt"

func main(){

	a := 42
	ptrA := &a

	fmt.Println("a : ", a, "ptrA : ", ptrA, "*ptrA", *ptrA)


}

/*
output :
a :  42 ptrA :  0xc042060080 *ptrA 42
*/