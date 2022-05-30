package main

import (
	"design_pattern/option"
)

func main() {
	p1 := option.New(20000, option.WithName("why"), option.WithAge(25))
	p1.PrintName()
	p1.PrintSalary()
	p1.PrintAge()

	p2 := option.New(15000, func(option *option.Option) {
		option.Name = "xp"
		option.Age = 36
	})
	p2.PrintName()
	p2.PrintSalary()
	p2.PrintAge()

	s1 := option.New2(option.WithName2("why"), option.WithAge2(25))
	s1.PrintName()
	s1.PrintAge()

	s2 := option.New2(func(student *option.Student) {
		student.Name = "xp"
		student.Age = 36
	})
	s2.PrintName()
	s2.PrintAge()
}
