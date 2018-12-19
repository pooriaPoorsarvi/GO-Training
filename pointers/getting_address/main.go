package main

import "fmt"

// everything is PASS BY VALUE in go, btw
// when we pass a memory address, we are passing a value

func main(){

	a := 42
	ptrA := &a

	fmt.Printf("%T \t %T \n", ptrA, a)
	fmt.Println("a : ", a, "ptrA : ", ptrA, "*ptrA", *ptrA)


}

/*
output :
*int 	 int
a :  42 ptrA :  0xc042060080 *ptrA 42
*/