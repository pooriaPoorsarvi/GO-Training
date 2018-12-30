package main

import "fmt"

type inter interface {

}


func main(){

	var a inter

	var a1 string


	a = "a"
	a1 = "a"

	// In The following line no conversion is done, we are just asserting,  assuming, that the underlying object is of type "string"
	fmt.Println(a.(string))
	// In the next line we actually convert our object, so we are making a new object
	a2 := []byte(a1)
	fmt.Printf("%v %p %p \n", a2, &a2, &a1)

	val1, ok1 := a.(string)
	// if we didn't use comma ok idiom we would have received a panic
	val2, ok2 := a.(int)

	if ok1 {
		fmt.Println(val1)
	}else {
		fmt.Println("Not a string")
	}
	if ok2 {
		fmt.Println(val2)
	}else {
		fmt.Println("Not an int")
	}



}



