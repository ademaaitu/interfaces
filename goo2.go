package main

import (
	"fmt"
	_ "math"
)

type Student interface {
	name_student() string
	number_group() int
	gpa() float64
}
type std struct {
	name             string
	group            int
	soc_gpa, edu_gpa float64
}

func (n std) name_student() string {
	return n.name
}
func (n std) number_group() int {
	return n.group
}
func (n std) gpa() float64 {
	return (n.soc_gpa + n.edu_gpa) / 2
}
func print(s Student) {
	fmt.Println("Name of studen:t", s.name_student())
	fmt.Println("Group number:", s.number_group())
	fmt.Println("GPA of student:", s.gpa())
}
func main() {
	n := std{
		name:    "John",
		group:   2106,
		soc_gpa: 3.5,
		edu_gpa: 3.7}
	print(n)
}
