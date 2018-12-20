package main

import "fmt"

func funcMaker() func() {
	return func() {
		fmt.Println("this function was made inside another function")
	}
}

func main(){

	madeFunc1 := func() {
		fmt.Println("this function was made inside the main function")
	}

	madeFunc2 := funcMaker()

	fmt.Printf("%T \n", madeFunc1)
	madeFunc1()

	fmt.Printf("%T \n", madeFunc2)
	madeFunc2()


}


