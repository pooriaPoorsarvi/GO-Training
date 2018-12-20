package main

import "fmt"

/*
function callback is the act of sending another function into another function as an argument
*/

func apply(funcToApply func (float64) float64, data ...float64) []float64{

	fmt.Printf("%T \n", data)
	for i, iv := range data{
		data[i] = funcToApply(iv)
	}
	return data
}


func main(){
	arr := [3]float64{1, 2, 3}
	fmt.Printf("%T \n", arr)
	funcToApply := func(a float64) float64{
		return a * a
	}
	fmt.Println(arr)
	apply(funcToApply, arr[:]...)
	fmt.Println(arr)
}

/*
output :
[3]float64
[1 2 3]
[]float64
[1 4 9]
*/


