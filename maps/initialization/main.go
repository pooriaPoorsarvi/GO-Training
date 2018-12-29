package main

import "fmt"

func main(){
	var map1 map[string]int
	map2 := new(map[string]int)
	//if you don't put the following line you will receive a panic because the map should not be nil
	map1 = make(map[string]int)
	map1["hey"] = 1
	fmt.Println(map1)
	//if you don't put the following line you will receive a panic because the map should not be nil
	*map2 = make(map[string]int)
	(*map2)["hey"] = 1
	fmt.Println(map2)
	map3 := map[string]int{
		"Pooria" : 20,
	}
	fmt.Println(map3)
}
/*
output :
map[hey:1]
&map[hey:1]
map[Pooria:20]
*/