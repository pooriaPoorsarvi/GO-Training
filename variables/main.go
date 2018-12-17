package main

import "fmt"

var A = 1
var B int
//C := 2


func main(){

	var D bool
	E := true


	var str string



	str = `You could have left this alone to have the zero value of "" :D`

	var a, b, c = 1, 2, 3
	d, e, f, g := 4, 5, 6, "even an string can be added here"

	fmt.Printf("%T \t %v \n", A, A)
	fmt.Printf("%T \t %v \n", B, B)
	fmt.Printf("%T \t %v \n", D, D)
	fmt.Printf("%T \t %v \n", E, E)
	fmt.Printf("%T \t %v \n", str, str)
	fmt.Printf("%T \t %v \n", a, a)
	fmt.Printf("%T \t %v \n", b, b)
	fmt.Printf("%T \t %v \n", c, c)
	fmt.Printf("%T \t %v \n", d, d)
	fmt.Printf("%T \t %v \n", e, e)
	fmt.Printf("%T \t %v \n", f, f)
	fmt.Printf("%T \t %v \n", g, g)




}
