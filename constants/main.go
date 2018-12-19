package main

import "fmt"

const a = "first const"
const(
	a1 = iota
	b1 = iota
	c1 = iota
)
const(
	a2 = iota
	b2
	c2
)



func main(){
	//const b := 1
	const c = 1



	fmt.Println("a", a)
	fmt.Println("c", c)
	fmt.Println("a1", a1)
	fmt.Println("b1", b1)
	fmt.Println("c1", c1)

	fmt.Println("a2", a2)
	fmt.Println("b2", b2)
	fmt.Println("c2", c2)

	const(
		a2 = iota
		b2
		c2
	)

	fmt.Println("a2", a2)
	fmt.Println("b2", b2)
	fmt.Println("c2", c2)


}

/*
output :
a first const
c 1
a1 0
b1 1
c1 2
a2 0
b2 1
c2 2
a2 0
b2 1
c2 2
*/

