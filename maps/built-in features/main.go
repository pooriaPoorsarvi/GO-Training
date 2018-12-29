package main

import "fmt"

func main(){
	mainMap := map[string]int{
		"first" : 1,
	}
	fmt.Println("The map is ", mainMap,"and the len of the map is : ", len(mainMap))
	mainMap["first"] = 2
	fmt.Println("The map is ", mainMap,"and the len of the map is : ", len(mainMap))
	delete(mainMap, "first")
	fmt.Println("The map is ", mainMap,"and the len of the map is : ", len(mainMap))
	// The following line is called comma ok idiom, if you don't check something exists, you can still delete it but nothing will change if it didn't exist before
	//val, exists := mainMap["first"]
	if val, exists := mainMap["first"]; exists{
		fmt.Println("the key was found and the value was : ", val)
	}else {
		fmt.Println("the key was not found")
	}
	for i := 90 ; i <= 100 ; i++ {
		mainMap[string(i)] = i;
	}
	for key, value := range mainMap{
		fmt.Println(key, value)
	}
}

/*
output :
The map is  map[first:1] and the len of the map is :  1
The map is  map[first:2] and the len of the map is :  1
The map is  map[] and the len of the map is :  0
the key was not found
] 93
Z 90
\ 92
_ 95
` 96
a 97
b 98
c 99
d 100
[ 91
^ 94
*/


