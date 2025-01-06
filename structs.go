package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name + "io"}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Rey", 20})
	fmt.Println(person{name: "ali", age: 30})
	fmt.Println(person{name: "red"})
	fmt.Println(&person{name: "reu", age: 22})
	fmt.Println(newPerson("rei"))

	t := newPerson("hal")
	fmt.Println(t.name)

	s := person{name: "son", age: 33}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age, s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
