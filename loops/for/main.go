package main

import "fmt"

func main(){

	for i := 2 ; i < 5 ; i ++{
		fmt.Println("i", i)
	}
	i := 2
	for ; i < 5 ; i ++{
		fmt.Println("i", i)
	}

	i = 2
	for ; i < 5 ; {
		fmt.Println("i", i)
		i++
	}

	// the last one acts like a while loop
	i = 2
	for  i < 5 {
		fmt.Println("i", i)
		i++
	}



}

/*
output :
i 2
i 3
i 4
i 2
i 3
i 4
i 2
i 3
i 4
i 2
i 3
i 4
*/