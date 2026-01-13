package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"studentManagementCLI/model"
)

var Students []model.Student

var NextID = 1

func AddStudents() {
	var name, grade, ageStr string
	var age int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Student Name: ")
	name, _ = reader.ReadString('\n')

	fmt.Print("Enter Student Age: ")
	ageStr, _ = reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("❌ Invalid age input. Please enter a valid number.")
		return
	}
	fmt.Print("Enter the Grade")
	grade, _ = reader.ReadString('\n')

	student := model.Student{
		ID:    NextID,
		Name:  strings.TrimSpace(name),
		Age:   age,
		Grade: strings.TrimSpace(grade),
	}
	Students = append(Students, student)
	NextID++
	fmt.Println("✅ Student added successfully")

}
func ViewStudents() {
	if len(Students) == 0 {
		fmt.Println("No students found.")
		return
	}
	fmt.Println("\n--- Student List ---")
	for key, student := range Students {
		fmt.Printf("%v Student  : %+v\n", key+1, student)
	}
}
func SearchStuent() {
	if len(Students) == 0 {
		fmt.Println("No students found.")
		return
	}

	studentMap := make(map[int]model.Student)
	for _, std := range Students {
		studentMap[std.ID] = std
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Student ID to search: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ Invalid ID input for search. Please enter a valid number.")
		return
	}
	if student, foundBool := studentMap[id]; foundBool {
		fmt.Printf("✅ Student found: %+v\n", student)
	} else {
		fmt.Println("❌ Student with the given ID not found.")
	}
}
func UpdateStudent() {
	if len(Students) == 0 {
		fmt.Println("No students found.")
		return
	}
	readder := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Student ID to update: ")
	idStr, _ := readder.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ Invalid ID input for update. Please enter a valid number.")
		return
	}
	for i := range Students {
		if Students[i].ID == id {
			fmt.Print("Enter new name (leave blank to keep current): ")
			name, _ := readder.ReadString('\n')
			name = strings.TrimSpace(name)
			if name != "" {
				Students[i].Name = name
			}
			fmt.Print("Enter new Age (leave blank to keep current): ")
			ageStr, _ := readder.ReadString('\n')
			ageStr = strings.TrimSpace(ageStr)
			if ageStr != "" {
				age, err := strconv.Atoi(ageStr)
				if err == nil {
					Students[i].Age = age
				}
			}
			fmt.Print("Enter new Grade (leave blank to keep current): ")
			grade, _ := readder.ReadString('\n')
			grade = strings.TrimSpace(grade)
			if grade != "" {
				Students[i].Grade = grade
			}
			fmt.Println("✅ Student updated successfully.")
			return
		}

	}
	fmt.Println("❌ Student with the given ID not found.")
}

func DeleteStudent() {
	if len(Students) == 0 {
		fmt.Println("No students found.")
		return
	}
	readder := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Student ID to update: ")
	idStr, _ := readder.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("❌ Invalid ID input for Delete . Please enter a valid number.")
		return
	}
	for i := range Students {
		if Students[i].ID == id {
			Students = append(Students[:i], Students[i+1:]...)
			fmt.Println("✅ Student deleted successfully.")
			return
		}
	}
	fmt.Println("❌ Student with the given ID not found.")

}

func SaveToFile() {

	file, err := os.Create("student.txt")
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, student := range Students {
		fmt.Fprintf(writer, "%d|%s|%d|%s\n", student.ID, student.Name, student.Age, student.Grade)
	}
	writer.Flush()
	fmt.Println("✅ Students saved to student.txt successfully.")
}

func LoadFromFile() {
	file, err := os.Open("student.txt")
	if err != nil {
		fmt.Println("No student data file found. Starting with empty student list.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")

		if len(parts) != 4 {
			fmt.Println("⚠️ Skipping invalid line:", line)
			continue
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("⚠️ Invalid ID in line:", line)
			continue
		}

		age, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("⚠️ Invalid age in line:", line)
			continue
		}

		student := model.Student{
			ID:    id,
			Name:  strings.TrimSpace(parts[1]),
			Age:   age,
			Grade: strings.TrimSpace(parts[3]),
		}
		Students = append(Students, student)
		if id >= NextID {
			NextID = id + 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("✅ Loaded %d students from file.\n", len(Students))
}
