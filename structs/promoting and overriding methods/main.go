package main

import "fmt"

type Father struct{
	lastName string
}

type Son struct {
	Father
	age int
}

type Daughter struct {
	Father
	name string
}


func (f Father) str() string{
	return f.lastName
}

func (s Son) str() string {
	return s.lastName + " " + fmt.Sprintf("%v", s.age)
}


func main(){
	s := Son{
		Father{
			"Son",
		},
		20,
	}
	d := Daughter{
		Father{
			"Daughter",
		},
		"Of her Father",
	}
	f := Father{
		"Father",
	}
	fmt.Println(s.str())
	fmt.Println(s.Father.str())
	fmt.Println(d.str())
	fmt.Println(f.str())
}
/*
output :
Son 20
Son
Daughter
Father
*/



