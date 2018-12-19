package main

import "fmt"

func main(){

	if true {
		fmt.Println("this will be printed")
	}else if true {
		fmt.Println("this won't be printed")
	}else {
		fmt.Println("this won't be printed")
	}

	if a:=true; a{
		fmt.Println("this will always be printed")
	}

	//the following line will cause an error because of the scope of a
	//fmt.Println(a)






}

