package main
import (
	"fmt"
	"time"
)

type Student struct {
	Email string
	Password string
}

type Teacher struct {
	Email string
	Password string
}

func TeacherWork(){
	var subject_count int
	var students_count int
	fmt.Printf("Enter number of subjects :")
	fmt.Scanln(&subject_count)
	fmt.Printf("Enter number of students :")
	fmt.Scanln(&students_count)
	var SubjectCount = make([]int, subject_count)
	var StudentArray = make([]int, students_count)
	var StudentPercentArray = make([]int, students_count)
	var sum_of_marks int
	for i := 0; i < students_count; i++ {
		fmt.Print("Enter student ", i+1, "ID :")
		fmt.Scanln(&StudentArray[i])
		for j := 0; j < subject_count; j++ {
			fmt.Print("Enter Subject ", j+1, " marks: ")
			fmt.Scanln(&SubjectCount[j])
			sum_of_marks = sum_of_marks + SubjectCount[i]
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
}

func main() {
	var start_time = time.Now()
	fmt.Printf("Hello Welcome to ABC Services : \n")
	fmt.Printf("1. Login as Student\n")
	fmt.Printf("2. Login as Teacher\nEnter 1 or 2 as input: ")
	var logininp int
	var Emailinp string
	var Passwordinp string
	var student1 Student
	var teacher1 Teacher
	student1.Email = "bob@gmail.com"
	student1.Password = "bob@123"
	teacher1.Email = "alice@gmail.com"
	teacher1.Password = "alice@123"
	fmt.Scanln(&logininp)
	
	if logininp == 1 {
		fmt.Printf("Enter Email: ")
		fmt.Scanln(&Emailinp)
		fmt.Printf("Enter Password: ")
		fmt.Scanln(&Passwordinp)
		if Emailinp == student1.Email && Passwordinp == student1.Password{
			fmt.Printf("\nSuccessfull Login")
		}else{
			fmt.Print("\nError in login. Please try again.")
		}
	}else if logininp == 2 {
		fmt.Printf("Enter Email: ")
		fmt.Scanln(&Emailinp)
		fmt.Printf("Enter Password: ")
		fmt.Scanln(&Passwordinp)
		if Emailinp == teacher1.Email && Passwordinp == teacher1.Password{
			fmt.Printf("\nSuccessfull login\n")
			TeacherWork()
		}else{
			fmt.Print("\nError in login. Please try again.")
		}
	}else{
		fmt.Printf("\nInvalid input.Please enter either 1 or 2")
	}


	elapsed := time.Since(start_time)
	fmt.Printf("\n%s", elapsed)
}
