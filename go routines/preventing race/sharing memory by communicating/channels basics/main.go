package main

import "fmt"

/*
we are gonna check unbuffered channels, you can also make buffered channels using make(chan T, m), and m will be the
size of your buffer
*/



func main(){
	// please keep in mind how we used b as a flag instead of using a wait group, so that we won't have to share memory
	//for the purpose of communicating, the use of this flag is also called semaphore
	c := make(chan int)
	// we have to use make in order to initialize a channel because it has underlying types that need to be made
	//in order to use it
	b := make(chan bool)


	go func() {
		for i := 0 ; i < 4 ; i ++ {
			c<-i
		}
		close(c)
	}()


	go func() {
		for i := range c{
			fmt.Println(i)
		}
		b<-true
	}()



	if <-b {
		fmt.Println("Program finished")
		close(b)
	}



}




/*
output :
0
1
2
3
Program finished
*/


