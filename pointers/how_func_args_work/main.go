package main

import "fmt"

// everything is PASS BY VALUE in go, btw
// when we pass a memory address, we are passing a value

func increment1(x int) int {
	x += 1
	return x
}

func increment2(x *int) int{
	v := *x
	v += 1
	return v
}

func increment3(x *int) int{
	*x += 1
	return *x
}

func main(){
	x := 0
	increment1(x)
	increment1(x)
	fmt.Println("the value of x is : ", x)
	increment2(&x)
	increment2(&x)
	fmt.Println("the value of x is : ", x)
	increment3(&x)
	increment3(&x)
	fmt.Println("the value of x is : ", x)

}

/*
output :
the value of x is :  0
the value of x is :  0
the value of x is :  2
*/
