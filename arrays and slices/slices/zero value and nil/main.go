package main

import "fmt"

/*
There is something that you need to know before you continue, and that is the difference between composite literals and new and make
you can read every thing we talk about in https://golang.org/doc/effective_go.html#allocation_new and in fact all of the following is
just a summary for the aforementioned link :D
"new" : "it does not initialize the memory, it only zeros it", so as it's said in the document we will have just acquired some space we can use
for our object and if we need to use some property of it, it is set to its zero value, we can change this, with one of two ways :
either use another function like make in order to change the value of *ptr to something we want or call a constructor for our ptr
keep in mind that by constructor we mean something like the following :
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
"composite literals" : this type of initialization were used to make local variables but since golang
decides which variables point towards stack and which ones point point towards the heap, &{}T and new(T) are alike but they are not always the same as is discussed in
https://github.com/golang/go/issues/29425 because when we use composite literals we are constructing our value, even if all the values are set to zero values.
for the properties are set to zero value in the second part. it's also easier to initialize default
values of properties using this method
"make" : "It creates slices, maps, and channels only" and "returns an initialized (not zeroed) value of type T (not *T).
The reason for the distinction is that these three types represent, under the covers, references to data structures that must be initialized before use."
so because we need underlying types in order to use something we will need to use make instead of new.
P.S make sure that you are not using index of nil values.
*/

func main(){



	var arr [] int
	fmt.Println(arr)
	fmt.Println(arr == nil)
	fmt.Println()
	arr1 := new([]int)
	fmt.Println(*arr1, len(*arr1))
	fmt.Println(*arr1 == nil)
	fmt.Println(arr1 == nil)
	fmt.Println()
	arr2 :=  &[]int {}
	fmt.Println(*arr2, len(*arr2))
	fmt.Println(*arr2 == nil)
	fmt.Println(arr2 == nil)
	fmt.Println()
	arr3 := make([]int, 10)
	fmt.Println(arr3 == nil)
	fmt.Println()
}
