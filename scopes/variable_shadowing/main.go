package main

import "fmt"

func shadow() string {
	return "this function is going to be shadowed, with a variable :D"
}

func main(){

	shadow := shadow()
	fmt.Println(shadow)

}
/*
output :
this function is going to be shadowed, with a variable :D
*/