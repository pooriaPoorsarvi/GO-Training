package main

import "fmt"

func main(){
	var a string
	a = "I am an example"

	fmt.Printf("The value is : %v and The type is : %T \n", a, a)

	fmt.Printf("The value is : %v and The type is : %T \n", a[0], a[0])

	fmt.Printf("The value is : %v and The type is : %T \n", rune(a[0]), rune(a[0]))


}
