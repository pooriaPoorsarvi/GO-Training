package main

import "fmt"


/*
Keep in mind that every argument in go is sent by value, but values them selves can be pointers, that is the case with
slices but that is not the case with arrays :D
*/

func change(a [2] int){
	fmt.Printf("%p %T \n",&a, a)
	fmt.Println(a)
	a[1] = 100;
	fmt.Println(a)
}


func main(){

	arr := [...]int{
		1,2,
	}
	fmt.Printf("%p %T \n", &arr, arr)
	fmt.Println(arr)
	change(arr)
	fmt.Println(arr)


}

/*
output :
0xc042062080 [2]int
[1 2]
0xc0420620d0 [2]int
[1 2]
[1 100]
[1 2]
*/