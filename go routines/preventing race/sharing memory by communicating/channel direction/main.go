package main

import (
	"runtime"
	"fmt"
)

func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}



func main(){


	c := incrementer()

	for i := range puller(c){
		fmt.Println(i)
	}

}


func incrementer() <- chan int{
	out := make(chan int)


	go func() {
		for i := 0 ; i < 5 ; i ++ {
			out <- i
		}
		// VERY IMPORTANT : even though this function is not the main function, if it doesn't close its channel
		//The channel in puller will not be closed and because of that the our main will never end
		//So this will end in to a DEAD LOCK so we need to have the following line
		// comment it out and see what happens :D
		close(out)
	}()


	return out
}

func puller(a <- chan int) <- chan int{
	out := make(chan int)

	var sum int

	go func() {

		for i := range a {
			sum += i
		}

		out <- sum
		close(out)
	}()

	return out


}

/*
output :
10
*/




