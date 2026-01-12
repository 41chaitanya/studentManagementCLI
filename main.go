package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"studentManagementCLI/service"
)
const (
	ADD_STUDENT    = 1
	VIEW_STUDENTS  = 2
	SEARCH_STUDENT = 3
	UPDATE_STUDENT = 4
	DELETE_STUDENT = 5
	EXIT           = 6
)
func main() {
	fmt.Println("+====================================+")
	fmt.Println("|    Student Management CLI Tool     |")
	fmt.Println("+====================================+")
	fmt.Println("Welcome ! You are inside the app.")

	for {
		fmt.Println("Menue ----- :)")
		fmt.Println("1. Add Student")
		fmt.Println("2. View Students")
		fmt.Println("3. Search Student")
		fmt.Println("4. Update Student")
		fmt.Println("5. Delete Student")
		fmt.Println("6. Exit")

		choice :=getUserChoice()
		if choice ==0{
			continue
		}
		if choice ==6{
			fmt.Println("Exiting the application. Goodbye!")
			break
		}
		choiceHandling(choice)
		
	}
}

	func choiceHandling(choice int){
		switch choice {
			case ADD_STUDENT:
				service.AddStudents()
			case VIEW_STUDENTS:
				service.ViewStudents()
			case SEARCH_STUDENT:
				service.SearchStuent()
			case UPDATE_STUDENT:
				service.UpdateStudent()
			case DELETE_STUDENT:
				service.DeleteStudent()
			default:
				fmt.Println("Invalid choice. Please try again.")

		}
		
	}
	func getUserChoice() int {
	
		reader :=bufio.NewReader(os.Stdin)
		fmt.Print("Enter your choice (1-6):")

		input ,_ :=reader.ReadString('\n')
		input = strings.TrimSpace(input)
		numInput,err:=strconv.Atoi(input)
		if err!=nil{
			fmt.Println("❌ Invalid input. Please enter a number. ONLY NUMBERS ARE ALLOWED.")
			return 0
		}
		if numInput <1 || numInput>6{
			fmt.Println("❌ Please enter a number between 1 and 6.")
			return 0
		}
		return numInput
	}
