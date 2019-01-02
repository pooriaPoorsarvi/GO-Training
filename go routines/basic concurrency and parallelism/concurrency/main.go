package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup



func main(){

	wg.Add(2)

	go funcCalling("Foo")
	go funcCalling("Bar")

	// if we don't add the wait group we will exit the main before the other two function calls end, so we will not see the outputs.
	wg.Wait()



}

func funcCalling(name string){

	for i := 0 ; i < 5 ; i ++ {
		fmt.Println(name, i)
	}

	wg.Done()

}






