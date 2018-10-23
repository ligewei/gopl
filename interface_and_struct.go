package main

import "fmt"

type I interface {
	Get() int
	Set(int)
	AnotherSet(int)
}

type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {
	s.Age = age
}

func (s S) AnotherSet(age int) {
	s.Age = age
}

func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())

	i.AnotherSet(100)
	fmt.Println(i.Get()) // 10
}

func main() {
	s := S{}
	// f(s) cannot use s (type S) as type I in argument to f, S does not implement I (Set method has pointer receiver)
	f(&s)
}
