package main

import "fmt"

/*
something very important that you should keep in mind that this variadic is infact a slice and you can only convert slices to this type of function argument
also if you do so, your data will be changes, because even though every thing is passed by value in golang, the actual value of a slice is a header which contains
a pointer to the actual array. more can be read at : https://stackoverflow.com/questions/39993688/are-golang-slices-pass-by-value
you can change your array to a slice using [:], but after doing this, your data may again change after passing it to a function
*/

func mean(numbers ...float64) float64{

	var total float64

	for _, v := range numbers{
		total += v
	}

	total /= float64(len(numbers))

	return total

}

func meanForSlice(numbers []float64) float64{

	var total float64

	for _, v := range numbers{
		total += v
	}

	total /= float64(len(numbers))

	return total

}

func main(){


	fmt.Println(mean(1, 2, 3))


	arr := []float64{1, 2, 3}

	fmt.Println(mean(arr...))

	fmt.Println(meanForSlice(arr))


}
/*
output :
2
2
2
*/