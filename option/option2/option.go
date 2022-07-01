package option2

import "fmt"

type human struct {
	name string
	age  int
	opt  opt
}

type opt struct {
	country string
	city    string
}

type OtherInfo interface {
	apply(*opt)
}

type funcOpt struct {
	fn func(*opt)
}

func (f *funcOpt) apply(do *opt) {
	f.fn(do)
}

func WithCountry(country string) OtherInfo {
	return &funcOpt{
		fn: func(o *opt) {
			o.country = country
		},
	}
}

func WithCity(city string) OtherInfo {
	return &funcOpt{
		fn: func(o *opt) {
			o.city = city
		},
	}
}

//type funcOpt func(*opt)
//
//func (f funcOpt) apply(do *opt) {
//	f(do)
//}
//
//func WithCountry(country string) OtherInfo {
//	return funcOpt(func(o *opt) {
//		o.country = country
//	})
//}
//
//func WithCity(city string) OtherInfo {
//	return funcOpt(func(o *opt) {
//		o.city = city
//	})
//}

func NewHuman(name string, age int, opt ...OtherInfo) *human {
	h := &human{
		name: name,
		age:  age,
		opt:  defaultOpt(),
	}

	for _, fn := range opt {
		fn.apply(&h.opt)
	}

	return h
}

func defaultOpt() opt {
	return opt{
		country: "China",
		city:    "Wuhan",
	}
}

func (h *human) Print() {
	fmt.Printf("%+v", *h)
}
