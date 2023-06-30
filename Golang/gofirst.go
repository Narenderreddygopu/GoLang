package main

import (
	"fmt"
)

func main() {

	var total_students_count int = 4
	var student1 string = "Narender" //type is string
	var student2 string = "Ramya"    //type is inferred
	var student3 string = "Uthej"
	var student4 string = "Harish"
	x := 4 //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(student3)
	fmt.Println(student4)
	fmt.Println(x)
	/*output*/
	fmt.Println("total_students_count", total_students_count)

}
