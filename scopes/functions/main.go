package main

import "fmt"

var x_o int

func increment1(x int) int {
	x += 1
	return x
}

func increment2() int {
	x_o += 1
	return x_o
}

func wrapper1() func() int {
	x := 0

	increment3 := func () int {
		x += 1
		return x
	}

	return increment3

}

func wrapper2(x int) func() int {


	/*
	this is called an anonymous function
	*/
	return func () int {
		x += 1
		return x
	}


}

func main(){
	x_in := 0
	increment1(x_in)
	out := increment1(x_in)
	fmt.Println("increment1 was called twice and the result was as the following : ", out)
	increment2()
	out = increment2()
	fmt.Println("increment2 was called twice and the result was as the following : ", out)
	increment3 := wrapper1()
	increment3()
	out = increment3()
	fmt.Println("increment3 was called twice and the result was as the following : ", out)
	x_in = 0
	increment4 := wrapper2(x_in)
	increment4()
	out = increment4()
	fmt.Println("increment4 was called twice and the result was as the following : ", out)

}


/*
output :
increment1 was called twice and the result was as the following :  1
increment2 was called twice and the result was as the following :  2
increment3 was called twice and the result was as the following :  2
increment4 was called twice and the result was as the following :  2
*/

