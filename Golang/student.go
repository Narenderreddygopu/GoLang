package main

import (
	"fmt"
	"time"
	"./mypackages"
)

type Student struct {
	Email    string
	Password string
}

type Teacher struct {
	Email    string
	Password string
}

var subject_count int
var students_count int
var student_marks int
var SubjectCount []int
var StudentArray []int
var StudentPercentArray []int
var sum_of_marks int


func TeacherWork() {

	fmt.Printf("Enter number of subjects :")
	fmt.Scanln(&subject_count)
	fmt.Printf("Enter number of students :")
	fmt.Scanln(&students_count)
	SubjectCount = make([]int, subject_count)
	StudentArray = make([]int, students_count)
	StudentPercentArray = make([]int, students_count)
	for i := 0; i < students_count; i++ {
		fmt.Print("Enter student ", i+1, " ID :")
		fmt.Scanln(&StudentArray[i])
		for j := 0; j < subject_count; j++ {
			fmt.Print("Enter Subject ", j+1, " marks: ")
			fmt.Scanln(&SubjectCount[j])
			sum_of_marks = sum_of_marks + SubjectCount[j]
		}
		StudentPercentArray[i] = sum_of_marks / subject_count
		sum_of_marks = 0
	}
	for i := 0; i < students_count; i++ {
		switch {
		case StudentPercentArray[i] >= 90:
			fmt.Print("\nstudent ", i+1, " secured a percentage of  :", StudentPercentArray[i], "% and received an A grade.")
		case StudentPercentArray[i] >= 80 && StudentPercentArray[i] < 90:
			fmt.Print("\nstudent ", i+1, " secured a percentage of  :", StudentPercentArray[i], "% and received an B grade.")
		case StudentPercentArray[i] >= 70 && StudentPercentArray[i] < 80:
			fmt.Print("\nstudent ", i+1, " secured a percentage of  :", StudentPercentArray[i], "% and received an C grade.")
		case StudentPercentArray[i] >= 60 && StudentPercentArray[i] < 70:
			fmt.Print("\nstudent ", i+1, " secured a percentage of  :", StudentPercentArray[i], "% and received an D grade.")
		case StudentPercentArray[i] >= 50 && StudentPercentArray[i] < 60:
			fmt.Print("\nstudent ", i+1, " secured a percentage of  :", StudentPercentArray[i], "% and received an E grade.")
		case StudentPercentArray[i] < 50:
			fmt.Print("\nstudent ", i+1, " secured a percentage of  :", StudentPercentArray[i], "% and received an F grade.")
		}
	}
	StudentWork()
}

func StudentWork() {
	var student_Id int
	var Id_Index int = 0
	var Recheck_Confirmation string
	for {
		fmt.Print("\nEnter Student_Id : \n")
		fmt.Scanln(&student_Id)
		for i := 0; i < len(StudentArray); i++ {
			if StudentArray[i] == student_Id {
				Id_Index = i + 1
				break
			}
		}
		if Id_Index != 0 {
			fmt.Print("\nYour score is :", StudentPercentArray[Id_Index-1])
		} else {
			fmt.Print("\nEnetered Wrong Student ID")
		}
		fmt.Println("\nEnter yes or no to recheck student percentage \n")
		fmt.Scanln(&Recheck_Confirmation)
		if Recheck_Confirmation == "No" {
			break
		}
	}
}

func main() {
	var start_time = time.Now()
	fmt.Printf("Hello Welcome to NR Services : \n")
	fmt.Printf("1. Login as Student\n")
	fmt.Printf("2. Login as Teacher\nEnter 1 or 2 as input: ")
	var logininp int
	var Emailinp string
	var Passwordinp string
	// var student1 Student
	// var teacher1 Teacher
	// student1.Email = "nr@gmail.com"
	// student1.Password = "nr@123"
	// teacher1.Email = "rn@gmail.com"
	// teacher1.Password = "rn@123"

	S1 := make(map[string]Student)
	T1 := make(map[string]Teacher)
	S1["Student_login"] = Student{"nr@gmail.com", "nr@123"}
	T1["Teacher_login"] = Teacher{"rn@gmail.com", "rn@123"}
	student1 := S1["Student_login"]
	teacher1 := T1["Teacher_login"]
	if logininp == 1 {
		fmt.Printf("Enter Email: ")
		fmt.Scanln(&Emailinp)
		fmt.Printf("Enter Password: ")
		fmt.Scanln(&Passwordinp)
		if myPackage.isEmailValid(Emailinp) && myPackage.isPasswardValid(Passwordinp) {
			if Emailinp == student1.Email && Passwordinp == student1.Password {
				fmt.Printf("\nSuccessfull Login")
			} else {
				fmt.Print("\nError in login. Please try again.")
			}
		} else {
			fmt.Println("not a valid email or passward")
		}

	} else if logininp == 2 {
		fmt.Printf("Enter Email: ")
		fmt.Scanln(&Emailinp)
		fmt.Printf("Enter Password: ")
		fmt.Scanln(&Passwordinp)
		if myPackage.isEmailValid(Emailinp) && myPackage.isPasswardValid(Passwordinp) {
			if Emailinp == teacher1.Email && Passwordinp == teacher1.Password {
				fmt.Printf("\nSuccessfull login\n")
				TeacherWork()
			} else {
				fmt.Printf("\nInvalid input.Please enter either 1 or 2")
			}
		} else {
			fmt.Println("not a valid email or passward")
		}
	}
	elapsed := time.Since(start_time)
	fmt.Printf("\n%s", elapsed)
}
