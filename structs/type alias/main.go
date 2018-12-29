package main

import "fmt"

//https://yourbasic.org/golang/type-alias/

type foo int
// The following line is similar to the way that runes are defined in golang
type bar = int

func main()  {
	var a foo
	var b bar


	a = 0
	b = 0

	c := b + 1

	c += int(a)
	//The following Line would have caused a problem
	//c += a

	fmt.Printf("Value : %v , Type : %T \n", a, a)
	fmt.Printf("Value : %v , Type : %T \n", b, b)
	fmt.Printf("Value : %v , Type : %T \n", c, c)


}


/*
output :
Value : 0 , Type : main.foo
Value : 0 , Type : int
Value : 1 , Type : int
*/