package main

import (
	"sync"
	"runtime"
	"fmt"
	"sync/atomic"
)

var wg sync.WaitGroup


//please take note that we couldn't just use int, we had to specify intX so that we could use atomic functions
var cnt int32

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
		atomic.AddInt32(&cnt, 1)
		fmt.Println("Out", name," value : ", cnt)


	}
	wg.Done()
}



/*
output:
In Bar  value :  0
Out Bar  value :  1
In Bar  value :  1
Out Bar  value :  2
In Bar  value :  2
Out Bar  value :  3
In Bar  value :  3
Out Bar  value :  4
In Bar  value :  4
Out Bar  value :  5
In Bar  value :  5
Out Bar  value :  6
In Bar  value :  6
Out Bar  value :  7
In Bar  value :  7
Out Bar  value :  8
In Bar  value :  8
Out Bar  value :  9
In Bar  value :  9
Out Bar  value :  10
In Bar  value :  10
Out Bar  value :  11
In Bar  value :  11
Out Bar  value :  12
In Bar  value :  12
Out Bar  value :  13
In Bar  value :  13
Out Bar  value :  14
In Bar  value :  14
Out Bar  value :  15
In Bar  value :  15
Out Bar  value :  16
In Bar  value :  16
Out Bar  value :  17
In Bar  value :  17
Out Bar  value :  18
In Bar  value :  18
Out Bar  value :  19
In Bar  value :  19
Out Bar  value :  20
In Foo  value :  20
Out Foo  value :  21
In Foo  value :  21
Out Foo  value :  22
In Foo  value :  22
Out Foo  value :  23
In Foo  value :  23
Out Foo  value :  24
In Foo  value :  24
Out Foo  value :  25
In Foo  value :  25
Out Foo  value :  26
In Foo  value :  26
Out Foo  value :  27
In Foo  value :  27
Out Foo  value :  28
In Foo  value :  28
Out Foo  value :  29
In Foo  value :  29
Out Foo  value :  30
In Foo  value :  30
Out Foo  value :  31
In Foo  value :  31
Out Foo  value :  32
In Foo  value :  32
Out Foo  value :  33
In Foo  value :  33
Out Foo  value :  34
In Foo  value :  34
Out Foo  value :  35
In Foo  value :  35
Out Foo  value :  36
In Foo  value :  36
Out Foo  value :  37
In Foo  value :  37
Out Foo  value :  38
In Foo  value :  38
Out Foo  value :  39
In Foo  value :  39
Out Foo  value :  40
Final Counter :  40
*/
