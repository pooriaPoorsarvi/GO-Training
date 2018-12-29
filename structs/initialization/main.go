package main

import "fmt"

type myStruct struct {
	name string
	age int
}


func main(){
	var a myStruct

	fmt.Println(&a, a)


	b := new(myStruct)
	fmt.Printf("%p \n", b)




	fmt.Println(&b, b, *b)

	fmt.Println(&a == b)
	fmt.Println(a == *b)

	d := b

	b = &myStruct{"Pooria", 20}
	d.name = "Pooria"
	d.age = 20

	fmt.Printf("%v , %p \n", &b, b)
	fmt.Printf("%v , %p \n", &d, d)
}
/*
output :
&{ 0} { 0}
0xc04205a460
0xc042080020 &{ 0} { 0}
false
true
0xc042080020 , 0xc04205a4c0
0xc042080028 , 0xc04205a460
*/

