package option

import "fmt"

type people struct {
	salary int
	option Option
}

type Option struct {
	Name string
	Age  int
}

type modOption func(option *Option)

func WithName(name string) modOption {
	return func(option *Option) {
		option.Name = name
	}
}

func WithAge(age int) modOption {
	return func(option *Option) {
		option.Age = age
	}
}

func New(salary int, modOption ...modOption) *people {
	option := Option{
		Name: "why",
		Age:  25,
	}

	for _, fn := range modOption {
		fn(&option)
	}

	return &people{
		salary: salary,
		option: option,
	}
}

func (p *people) PrintSalary() {
	fmt.Println(p.salary)
}

func (p *people) PrintName() {
	fmt.Println(p.option.Name)
}

func (p *people) PrintAge() {
	fmt.Println(p.option.Age)
}

// ---------------------------------------------------------------------

type Student struct {
	Name string
	Age  int
}

type option func(student *Student)

func WithName2(name string) option {
	return func(student *Student) {
		student.Name = name
	}
}

func WithAge2(age int) option {
	return func(student *Student) {
		student.Age = age
	}
}

func New2(options ...option) Student {
	student := Student{}

	for _, fn := range options {
		fn(&student)
	}

	return student
}

func (s *Student) PrintName() {
	fmt.Println(s.Name)
}

func (s *Student) PrintAge() {
	fmt.Println(s.Age)
}

//func main() {
//	p1 := option.New(20000, option.WithName("why"), option.WithAge(25))
//	p1.PrintName()
//	p1.PrintSalary()
//	p1.PrintAge()
//
//	p2 := option.New(15000, func(option *option.Option) {
//		option.Name = "xp"
//		option.Age = 36
//	})
//	p2.PrintName()
//	p2.PrintSalary()
//	p2.PrintAge()
//
//	s1 := option.New2(option.WithName2("why"), option.WithAge2(25))
//	s1.PrintName()
//	s1.PrintAge()
//
//	s2 := option.New2(func(student *option.Student) {
//		student.Name = "xp"
//		student.Age = 36
//	})
//	s2.PrintName()
//	s2.PrintAge()
//}
