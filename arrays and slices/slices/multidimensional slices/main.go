package main

import "fmt"

/*
Keep in mind the differences of append and adding thing manually by setting values into specific indexes
*/

func main(){
	student := make([]string, 35)
	students := make([][]string, 35)
	student[10] = "Pooria"
	fmt.Println(student, cap(student))
	student1 := append(student, "Poorsarvi")
	fmt.Println(student1, cap(student1))
	students[0] = student
	students = append(students, student1)
	fmt.Println(students)
}
/*
output :
[          Pooria                        ] 35
[          Pooria                         Poorsarvi] 72
[[          Pooria                        ] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [          Pooria                         Poorsarvi]]
*/