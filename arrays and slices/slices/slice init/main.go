package main

import "fmt"

/*
you can find a very interesting talk on different types of memory allocations in golang, in https://www.youtube.com/watch?v=zjoieOpy5hE
T{} or Composite literals are used to make local variables in contrast to new but since golang decides which variables go to stack
and which go to the heap, you can use &T{} and it will return the same thing.
*/

func main(){
	var arr [3]int
	s1 := arr[:]
	// the next line is type, length which is equal to capacity
	s2 := make([]int, 3)
	// the next line is type, length and capacity
	s3 := make([]int, 3, 5)
	s4 := new([]int)
	s5 := []int{}
	s6 := &[]int{}
	fmt.Printf("s1 type : %T , value : %v \n", s1, s1)
	fmt.Printf("s2 type : %T , value : %v \n", s2, s2)
	fmt.Printf("s3 type : %T , value : %v \n", s3, s3)
	fmt.Printf("s4 type : %T , value : %v \n", s4, s4)
	fmt.Printf("s5 type : %T , value : %v \n", s5, s5)
	fmt.Printf("s6 type : %T , value : %v \n", s6, s6)
}
/*
output :
s1 type : []int , value : [0 0 0]
s2 type : []int , value : [0 0 0]
s3 type : []int , value : [0 0 0]
s4 type : *[]int , value : &[]
s5 type : []int , value : []
s6 type : *[]int , value : &[]
*/

