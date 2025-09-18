package main

import (
	"math/rand"
	"sort"
)

type Student struct {
	Id     int
	Name   string
	Grades [5]int
}

type StudentWithAvg struct {
	Student Student
	Average float64
}

type StudentRawData struct {
	Name  string
	Marks [5]int
}
type Class struct {
	students map[int]Student
}

func InitializeClassroom() *Class {
	return &Class{}
}

func (c *Class) AddNewStudent(s StudentRawData) Student {
	newId := rand.Intn(900000) + 100000
	student := Student{newId, s.Name, s.Marks}

	if c.students == nil {
		c.students = make(map[int]Student)
	}

	c.students[newId] = student
	return student
}

func (c *Class) AddMultipleStudents(d []StudentRawData) {
	for _, s := range d {
		c.AddNewStudent(s)
	}
}

func (s *Student) AverageGrade() (avg float64) {
	for _, mark := range s.Grades {
		avg += float64(mark)
	}

	avg /= float64(len(s.Grades))
	return
}

func (c *Class) GetTopStudents(limit int) []StudentWithAvg {
	results := []StudentWithAvg{}

	for _, s := range c.students {
		avg := s.AverageGrade()
		results = append(results, StudentWithAvg{s, avg})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Average > results[j].Average
	})

	topStudents := results
	if len(results) > limit {
		topStudents = results[:limit]
	}

	return topStudents
}
