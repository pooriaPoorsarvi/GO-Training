package main

import (
	"fmt"
	"runtime"
)

/*
Here we are gonna use n channels that act as our inputs and m channels that read from these inputs
we manage these channels and close them before the program finishes
*/



func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}



func main(){

	var n int
	var m int

	c := make(chan int)
	inDone := make(chan bool)
	b := make(chan bool)

	n = 2
	m = 3


	for i := 0 ; i < n ; i ++ {
		/*
		Here we make our inputs
		*/
		go func() {
			for j := 0 ; j < 5 ; j ++{
				c <- j
			}
			//The following line can not be used because we may have other precesses that may want to use this channel
			//close(c)
			inDone <- true
		}()
	}

	go func() {
		/*
		Here we make sure that all of the channels that give us our inputs are closed
		and then we close the input channel
		*/
		for i := 0 ; i < n ; i ++ {
			if <- inDone{
				fmt.Println("Input Closed .")
			}else {
				fmt.Println("This shouldn't have been printed .")
			}
		}
		close(inDone)
		close(c)
	}()

	for i := 0 ; i < m ; i ++ {
		/*
		Here we make the functions that makes the functions that get the input and print them for us
		*/
		go func() {
			for j := range c {
				fmt.Println(j)
			}
			b<-true
		}()
	}


	for i := 0 ; i < m ; i ++ {
		/*
		Here we make sure that all of our functions that read the inputs are closed so we can close are
		so that we can close our channel that acts as a semaphore
		*/

		if <-b{
			fmt.Println("Reading Channel Closed .")
		}else {
			fmt.Println("This shouldn't have been printed .")
		}
		//This was my first bug when writing this :D, can you guess what I did wrong ?
		//close(b)
		//fmt.Println("Program was Finished")
	}

	close(b)
	fmt.Println("Program was Finished")

}

/*
output :
0
0
3
1
4
2
3
4
Input Closed .
2
Input Closed .
Reading Channel Closed .
1
Reading Channel Closed .
Reading Channel Closed .
Program was Finished
*/

