package main

import (
	"sync"
	"runtime"
	"fmt"
)

var wg sync.WaitGroup


func init(){
	//this function will be called before the main and here we try to use all the processors so that we get parallelism
	runtime.GOMAXPROCS(runtime.NumCPU())
}



func main(){

	wg.Add(2)

	go funcCalling("Foo")
	go funcCalling("Bar")

	wg.Wait()




}


func funcCalling(name string){

	for i := 0 ; i < 5 ; i ++ {
		fmt.Println(name, i)
	}

	wg.Done()

}









