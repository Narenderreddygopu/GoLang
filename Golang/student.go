package main

import (
	"fmt"
	"time"
)

func main() {
	var start_time = time.Now()
	var total_students_count int = 4

	var narender int

	var student2 int
	var student3 int
	var student4 int
	var avg_marks int
	var sum_of_marks int
	x := 4 //type is inferred
	elapsed := time.Since(start_time)
	fmt.Print("Enter narender marks: ")
	fmt.Scanln(&narender)

	fmt.Print("Enter student2 marks: ")
	fmt.Scanln(&student2)

	fmt.Print("Enter student3 marks: ")
	fmt.Scanln(&student3)

	fmt.Print("Enter student4 marks: ")
	fmt.Scanln(&student4)

	sum_of_marks = (narender + student2 + student3 + student4)
	avg_marks = (sum_of_marks) / 4

	fmt.Println(narender)
	fmt.Println(student2)
	fmt.Println(student3)
	fmt.Println(student4)
	fmt.Println(x)

	fmt.Printf("%d + %d + %d + %d = %d\n", narender, student2, student3, student4, avg_marks)

	fmt.Println("Avg marks of students : ", avg_marks)

	/*output*/
	fmt.Println("total_students_count", total_students_count)

	fmt.Printf("%s", elapsed)

}
