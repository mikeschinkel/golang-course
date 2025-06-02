package main

import (
	"fmt"
)

type StringSet struct {
	a *[1024]int
	b int
}

// type States [50]string
type States []string

func (s *StringSet) Set(ss StringSet) {
	*s = ss
}
func (s *StringSet) Concat() string {
	return s.a + *s.b
}

func (s *StringSet) Assign(val *string) {
	(*s).a = *val
	s.b = new(string)
	// Assign to value being pointed to
	*s.b = *val //Value of
	s.b = val   // Address of
	*val = "1"
}

type StringSlice []string

func (s StringSlice) Count() int {
	return len(s)
}
func main() {
	//s := make([]string, 10)
	s2 := new(StringSet)
	s3 := &StringSet{}
	//s := make(StringSlice, 100)
	//s := make(StringSlice, 0, 100)
	st := States{"GA", "PA"}
	s := make(StringSlice, len(st)) //Good
	for i, e := range st {
		s[i] = e
	}

	s = make(StringSlice, 0, 100) //Good
	for _, e := range st {
		s = append(s, e)
	}

	s = make(StringSlice, 0) //Bad
	for _, e := range st {
		s = append(s, e)
	}

	fmt.Println(s[0])
	s = append(s, "10")
	fmt.Println(s[100])
	fmt.Println(len(s))
	fmt.Println(s.Count(), s2.Concat())

	ss := StringSet{}
	println(ss.b)
	//ss.Assign("1")
	//ss.Assign("2")
	//x := "3"
	var x string
	var y *string
	ss.Assign(&x)
	println(ss.a)
	println(ss.b)
	println(x)
	println(y)
}
