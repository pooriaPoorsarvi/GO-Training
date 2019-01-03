package main

import (
	"fmt"
	"sync"
	"runtime"
	"time"
	"math/rand"
)


var wg sync.WaitGroup

var cnt int

var mutex sync.Mutex



func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}



func main(){

	wg.Add(2)


	go incrementer("Foo")
	go incrementer("Bar")

	wg.Wait()

	fmt.Println("Final Counter : ", cnt)




}



func incrementer(name string){

	for i := 0 ; i < 20 ; i ++ {

		fmt.Println("In", name," value : ", cnt)

		mutex.Lock()

		x := cnt

		x ++

		time.Sleep(time.Duration(rand.Intn(3)))

		cnt = x
		mutex.Unlock()
		fmt.Println("Out", name," value : ", cnt)

	}
	wg.Done()
}




/*
output :
In Foo  value :  0
In Bar  value :  0
Out Foo  value :  1
In Foo  value :  2
Out Bar  value :  2
In Bar  value :  2
Out Foo  value :  3
In Foo  value :  3
Out Bar  value :  4
In Bar  value :  4
Out Foo  value :  5
In Foo  value :  6
Out Bar  value :  6
In Bar  value :  6
Out Foo  value :  7
In Foo  value :  7
Out Bar  value :  8
In Bar  value :  8
Out Foo  value :  9
In Foo  value :  9
Out Foo  value :  10
In Foo  value :  10
Out Foo  value :  11
In Foo  value :  11
Out Bar  value :  12
In Bar  value :  12
Out Bar  value :  13
In Bar  value :  13
Out Bar  value :  14
In Bar  value :  14
Out Foo  value :  15
In Foo  value :  15
Out Bar  value :  16
In Bar  value :  17
Out Foo  value :  17
In Foo  value :  17
Out Bar  value :  18
In Bar  value :  18
Out Foo  value :  19
In Foo  value :  19
Out Bar  value :  20
In Bar  value :  20
Out Foo  value :  21
In Foo  value :  22
Out Bar  value :  22
In Bar  value :  22
Out Foo  value :  23
In Foo  value :  23
Out Bar  value :  24
In Bar  value :  25
Out Foo  value :  25
In Foo  value :  25
Out Bar  value :  26
In Bar  value :  26
Out Foo  value :  27
In Foo  value :  27
Out Bar  value :  28
In Bar  value :  28
Out Bar  value :  30
In Bar  value :  30
Out Foo  value :  29
In Foo  value :  30
Out Bar  value :  31
In Bar  value :  32
Out Foo  value :  32
In Foo  value :  32
Out Bar  value :  33
In Bar  value :  33
Out Foo  value :  34
In Foo  value :  34
Out Bar  value :  35
In Bar  value :  35
Out Foo  value :  36
In Foo  value :  36
Out Bar  value :  37
In Bar  value :  37
Out Foo  value :  38
In Foo  value :  38
Out Foo  value :  40
Out Bar  value :  39
Final Counter :  40
*/

