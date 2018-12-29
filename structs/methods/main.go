package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) fullName() string {
	return p.first + p.last
}

func main(){
	p1 := person{"Pooria", "Poorsarvi Tehrani", 20}
	fmt.Println(p1.fullName())
}


