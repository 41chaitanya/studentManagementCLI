package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		if choice < 1 || choice > 6{
			continue
		}
		if choice ==0{
			continue
		}
		fmt.Printf("nice you have entered %v \n\n",choice)
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
