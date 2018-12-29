package main

import "fmt"

type Father struct{
	a int
	b int
}
type Son struct {
	Father
	a int
}



func main(){
	bach := Son{}
	fmt.Println(bach)
	bach.a = 10
	bach.b = 20
	fmt.Println(bach)
	bach.Father.b = 30
	fmt.Println(bach)

}
/*
output :
{{0 0} 0}
{{0 20} 10}
{{0 30} 10}
*/






