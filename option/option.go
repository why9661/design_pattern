package main

import "fmt"

type People struct {
	salary int
	option Option
}

type Option struct {
	Name string
	Age  int
}

type ModOption func(option *Option)

func WithName(name string) ModOption {
	return func(option *Option) {
		option.Name = name
	}
}

func WithAge(age int) ModOption {
	return func(option *Option) {
		option.Age = age
	}
}

func New(salary int, modOption ...ModOption) *People {
	option := Option{
		Name: "why",
		Age:  25,
	}

	for _, fn := range modOption {
		fn(&option)
	}

	return &People{
		salary: salary,
		option: option,
	}
}

func main() {
	p := New(20000, WithName("why"))

	fmt.Println(p)
}
