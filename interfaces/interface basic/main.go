package main

import (
	"math"
	"fmt"
)

type shape interface {

	area () float64

}



type circle struct {

	R float64

}

type Square struct {

	Edge float64

}

func show(s shape){
	fmt.Println(s.area())
}

func (s Square) area() float64{
	return s.Edge * s.Edge
}

func (c circle) area() float64{
	return c.R * c.R * math.Pi
}





func main(){

	c := circle{1}
	s := Square{2}

	show(c)
	show(s)

}




