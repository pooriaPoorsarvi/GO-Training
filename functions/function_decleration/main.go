package main

import "fmt"

func noReturn(){
	fmt.Println("function with no return was called :D")
}

func withNormalReturn() string{
	return "function with normal return was called"
}

func withGoDecForReturn() (a string, b string){
	//Keep in mind that the following line would have caused a problem because a is already declared in the signature of the function :D and you can not shadow it because you need to return it
	//a := "function with golang return "
	a = "function with golang return "
	b = "was called"
	return
	//the following line would have caused no errors :D
	//return a, b
}


func main(){
	noReturn()
	fmt.Println(withNormalReturn())
	fmt.Println(withGoDecForReturn())
}

/*
output :
function with no return was called :D
function with normal return was called
function with golang return  was called
*/



