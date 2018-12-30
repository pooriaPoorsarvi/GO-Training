package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}


type Circle struct {

	R float64

}

type Square struct {

	Edge float64

}

func show(s Shape){
	fmt.Println(s.Area())
}

func (s Square) Area() float64{
	return s.Edge * s.Edge
}

func (c *Circle) Area() float64{
	return c.R * c.R * math.Pi
}



func main(){
	c1 := Circle{1}
	s1 := Square{2}


	c2 := &Circle{1}
	s2 := &Square{2}

	fmt.Println(c1, s1, c2, s2)

	//The following line will cause an error
	//show(c1)

	show(s1)

	show(c2)
	show(s2)

}



