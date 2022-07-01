package option1

import "fmt"

type people struct {
	salary int
	opt    Option
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

func NewPeople(salary int, modOption ...modOption) *people {
	p := &people{
		salary: salary,
		opt: Option{
			Name: "why",
			Age:  23,
		},
	}

	for _, fn := range modOption {
		fn(&p.opt)
	}

	return p
}

func (p *people) PrintSalary() {
	fmt.Println(p.salary)
}

func (p *people) PrintName() {
	fmt.Println(p.opt.Name)
}

func (p *people) PrintAge() {
	fmt.Println(p.opt.Age)
}
