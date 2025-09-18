package main

import "fmt"

func RunGradeManagement() {
	studentsData := []StudentRawData{
		{Name: "Aarav Mehta", Marks: []int{85, 90, 78, 92, 88}},
		{Name: "Isha Patel", Marks: []int{72, 80, 68, 75, 70}},
		{Name: "Rohan Sharma", Marks: []int{95, 92, 96, 94, 98}},
		{Name: "Neha Verma", Marks: []int{88, 84, 86, 90, 85}},
		{Name: "Kabir Iyer", Marks: []int{60, 65, 58, 62, 61}},
		{Name: "Simran Kaur", Marks: []int{78, 82, 80, 76, 79}},
		{Name: "Vikram Singh", Marks: []int{91, 89, 93, 90, 94}},
		{Name: "Ananya Rao", Marks: []int{83, 85, 88, 84, 82}},
		{Name: "Kunal Joshi", Marks: []int{55, 60, 52, 58, 57}},
		{Name: "Priya Nair", Marks: []int{87, 90, 85, 89, 88}},
	}

	class := InitializeClassroom()

	class.AddMultipleStudents(studentsData)
	class.AddNewStudent(StudentRawData{"Ketan Naik", []int{12, 34, 56, 78, 86}})

	fmt.Println("students are", class.students)
}
