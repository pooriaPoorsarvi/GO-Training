package main

import (
	"fmt"
	"sort"
	"sync"
	"runtime"
	"time"
)
var wg sync.WaitGroup

func foo(){
	for i:=1 ; i <40 ; i++ {
		fmt.Println("Foo:",i)
	}
	wg.Done()
}
func bar(){
	for i:=1 ; i <40 ; i++ {
		fmt.Println("Bar:",i)
	}
	wg.Done()
}

func test(a chan int) {
	fmt.Println(<-a)
}
func main() {
	//wg.Add(2)
	//go foo()
	//go bar()
	//wg.Wait()
	var a chan int
	//go func() {
	//	fmt.Println("hey	")
	//	b := <-a
	//	fmt.Println(b)
	//	close(a)
	//}()
	go test(a)
	time.Sleep(time.Second)
	var b int
	b = 2
	a <- b


}

func init()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}




type people []string
type pouya struct {
	people
}

func (anghezi people) Len() int {
	return len(anghezi)
}

func (this people) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this people) Swap(i, j int) {
	temp := this[i]
	this[i] = this[j]
	this[j] = temp
}







func firstExample()  {

	studyyGroup := people{"Zeno", "John", "Al", "Jenny"}

	s := []string{"Zeno", "John", "Al", "Jenny"}

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println(studyyGroup)

	fmt.Println(s)

	fmt.Println(n)

	sort.Sort(studyyGroup)

	sort.Strings(s)

	sort.Ints(n)

	fmt.Println(studyyGroup)

	fmt.Println(s)

	fmt.Println(n)


}
