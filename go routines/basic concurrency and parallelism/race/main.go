package main

import (
	"sync"
	"runtime"
	"time"
	"math/rand"
	"fmt"
)

// This code is an example of a race in a code because two different processes are trying to over write the same value which is cnt
// You can check out races during runtime by using : go run -race main.go instead of using the normal : go run main.go

var wg sync.WaitGroup


var cnt int


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

		x := cnt

		x ++

		time.Sleep(time.Duration(rand.Intn(3)))

		cnt = x

		fmt.Println("Out", name," value : ", cnt)

	}
	wg.Done()
}


/*
output :
In Foo  value :  0
In Bar  value :  0
Out Bar  value :  1
In Bar  value :  1
Out Foo  value :  1
In Foo  value :  2
Out Bar  value :  2
In Bar  value :  2
Out Foo  value :  3
In Foo  value :  3
Out Foo  value :  4
In Foo  value :  4
Out Bar  value :  3
In Bar  value :  4
Out Bar  value :  5
In Bar  value :  5
Out Foo  value :  5
In Foo  value :  5
Out Foo  value :  6
In Foo  value :  6
Out Bar  value :  6
In Bar  value :  6
Out Foo  value :  7
In Foo  value :  7
Out Foo  value :  8
In Foo  value :  8
Out Foo  value :  9
In Foo  value :  9
Out Bar  value :  9
In Bar  value :  9
Out Bar  value :  10
In Bar  value :  10
Out Foo  value :  10
In Foo  value :  11
Out Bar  value :  11
In Bar  value :  11
Out Foo  value :  12
In Foo  value :  12
Out Bar  value :  12
In Bar  value :  12
Out Foo  value :  13
In Foo  value :  13
Out Foo  value :  14
In Foo  value :  14
Out Bar  value :  13
In Bar  value :  14
Out Bar  value :  15
In Bar  value :  15
Out Foo  value :  15
In Foo  value :  16
Out Bar  value :  16
In Bar  value :  16
Out Bar  value :  17
In Bar  value :  17
Out Foo  value :  17
In Foo  value :  17
Out Foo  value :  18
In Foo  value :  18
Out Foo  value :  19
In Foo  value :  19
Out Bar  value :  18
In Bar  value :  20
Out Bar  value :  21
In Bar  value :  21
Out Foo  value :  20
In Foo  value :  21
Out Foo  value :  22
In Foo  value :  22
Out Foo  value :  23
In Foo  value :  23
Out Bar  value :  22
In Bar  value :  23
Out Bar  value :  24
In Bar  value :  24
Out Bar  value :  25
In Bar  value :  25
Out Bar  value :  26
In Bar  value :  26
Out Foo  value :  26
Out Bar  value :  27
Final Counter :  27
*/