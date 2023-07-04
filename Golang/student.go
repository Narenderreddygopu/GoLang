package main

import (
	"fmt"
	"time"
)

func main() {
	var start_time = time.Now()
	var subject_count int
	// var student_userId int
	var students_count int
	// fmt.Print("Enter student Id :")
	// fmt.Scanln(&Student_userId)
	fmt.Printf("Enter number of subjects :")
	fmt.Scanln(&subject_count)
	fmt.Printf("Enter number of students :")
	fmt.Scanln(&students_count)
	var array1 = make([]int, subject_count)
	var array2 = make([]int, students_count)
	var array3 = make([]int, students_count)
	var sum_of_marks int
	//var avg_marks int
	for i := 0; i < students_count; i++ {
		fmt.Print("Enter student ", i+1, "ID :")
		fmt.Scanln(&array2[i])
		for j := 0; j < subject_count; j++ {

			fmt.Print("Enter Subject ", j+1, " marks: ")
			fmt.Scanln(&array1[j])
			sum_of_marks = sum_of_marks + array1[i]
		}
		array3[i] = sum_of_marks / subject_count
		sum_of_marks = 0
	}
	// avg_marks = (sum_of_marks) / subject_count
	for i := 0; i < students_count; i++ {

		switch {
		case array3[i] >= 90:
			fmt.Print("\nstudent ", i+1, " secured % is  :", array3[i], " with A Garde ")
		case array3[i] >= 80 && array3[i] < 90:
			fmt.Print("\nstudent ", i+1, " secured  is  :", array3[i], " with B Garde ")
		case array3[i] >= 70 && array3[i] < 80:
			fmt.Print("\nstudent ", i+1, " secured % is  :", array3[i], " with C Garde ")
		case array3[i] >= 60 && array3[i] < 70:
			fmt.Print("\nstudent ", i+1, " secured % is  :", array3[i], " with D Garde ")
		case array3[i] >= 50 && array3[i] < 60:
			fmt.Print("\nstudent ", i+1, " secured % is  :", array3[i], " with E Garde ")
		case array3[i] < 50:
			fmt.Print("\nstudent ", i+1, " secured % is  :", array3[i], " with F Garde ")
		}
	}
	elapsed := time.Since(start_time)
	fmt.Printf("\n%s", elapsed)
}
