package main

import (
	"fmt"
	"time"
)

func main() {
	var start_time = time.Now()
	var array1 = [5]int{}
	// var total_students_count int = 4
	// var narender int
	// var student2 int
	// var student3 int
	// var student4 int
	var avg_marks int
	var sum_of_marks int
	// x := 4 //type is inferred

	// fmt.Print("Enter narender1 marks: ")
	// fmt.Scanln(&narender)

	// fmt.Print("Enter student2 marks: ")
	// fmt.Scanln(&student2)

	// fmt.Print("Enter student3 marks: ")
	// fmt.Scanln(&student3)

	// fmt.Print("Enter student4 marks: ")
	// fmt.Scanln(&student4)

	for i := 1; i < 5; i++ {
		fmt.Print("Enter Student", i, " marks: ")
		fmt.Scanln(&array1[i])
		sum_of_marks = sum_of_marks + array1[i]
	}

	// sum_of_marks = (narender + student2 + student3 + student4)
	avg_marks = (sum_of_marks) / 4
	elapsed := time.Since(start_time)
	// fmt.Println(narender)
	// fmt.Println(student2)
	// fmt.Println(student3)
	// fmt.Println(student4)
	// fmt.Println(x)

	// fmt.Printf("%d + %d + %d + %d = %d\n", narender, student2, student3, student4, avg_marks)

	fmt.Println("Avg marks of students : ", avg_marks)

	/*output*/
	// fmt.Println("total_students_count", total_students_count)

	fmt.Printf("%s", elapsed)

}
