package main

import "fmt"

func print(a [3]int, s []int)  {
	fmt.Printf("slice : %v array : %v size of slice : %v and size of array : %v \n", s, a, len(s), len(a))
}
func main(){

	var a [3]int
	s := a[:2]
	print(a, s)
	s2 := append(s, 1)
	print(a, s)
	print(a, s2)
	s2 = append(s2, 1)
	print(a, s2)
	s2 = append(s2, 1)
	print(a, s2)
	s2 = append(s2, 1)
	print(a, s2)
	s2 = append(s2, []int{1,2}...)
	print(a, s2)
	//The next line deletes the third element or the element at index 2
	s2 = append(s2[:2], s2[3:]...)
	print(a, s2)

}




