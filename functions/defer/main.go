package main

import "fmt"

func main(){

	defer fmt.Println("World !")

	fmt.Print("Hey ")


}
/*
output :
Hey World !
*/

