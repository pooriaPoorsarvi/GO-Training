package main

import "fmt"

func main(){



	funcCalling("Foo")
	funcCalling("Bar")

	fmt.Println("The above didn't use any kind of concurrency and parallelism \n")

}



func funcCalling(name string){

	for i := 0 ; i < 5 ; i ++ {
		fmt.Println(name, i)
	}


}


