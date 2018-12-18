package main

import "fmt"

func main(){
	x := 42


	{
		fmt.Println("This was called inside the block x ", x)

		y :=  "this is inside the block and can not be accessed from outside :D"
		fmt.Println("This was called inside the block y ", y)

		x :=  "this is The second x that was defined inside the block"
		fmt.Println("This was called inside the block x ", x)




	}
	fmt.Println("This was called outside the block x ", x)


}
