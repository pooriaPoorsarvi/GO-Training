package main

import (
	"fmt"
	"math/rand"
)

func main(){

	switch "condition" {
	case "condition":
		fmt.Println("condition was found")
		fmt.Println("keep in mind that we didn't have to add a block {} and also other conditions were not read")
	case "other":
		fmt.Println("this should not be printed !")
	default:
		fmt.Println("this also should not be printed !")
	}

	fmt.Println()

	switch "condition" {
	case "condition":
		fmt.Println("condition was found")
		fmt.Println("keep in mind that we didn't have to add a block {} and also other conditions were not read")
		fallthrough
	case "other":
		fmt.Println("this should be printed !")
		break
	default:
		fmt.Println("this should not be printed !")
	}

	fmt.Println()

	randN := rand.Int31n(10)

	var variable string

	if randN > 5 {
		variable = "ask"
	}else {
		variable = "question"
	}


	switch variable {

	case "ask", "question":
		fmt.Println("this will always be printed")
	default:
		fmt.Println("this will never be read")

	}

	fmt.Println()

	switch  {
	case true:
		fmt.Println("this will always be printed")
	default:
		fmt.Println("this will never be read")

	}

	fmt.Println()

	// here we are adding an abstraction layer to our variable :D
	var1 := interface{} (variable)

	//keep in mind that variable.(type) can only be used
	switch var1.(type) {
	case string:
		fmt.Println("this will always be printed")
	default:
		fmt.Println("this will never be read")
	}


}

/*
output :
condition was found
keep in mind that we didn't have to add a block {} and also other conditions were not read

condition was found
keep in mind that we didn't have to add a block {} and also other conditions were not read
this should be printed !

this will always be printed

this will always be printed

this will always be printed
*/

