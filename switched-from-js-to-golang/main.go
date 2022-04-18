package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

type Password string

func (p Password) validatePassword() error {
	if len(p) < 5 {
		return errors.New("password is too short")
	}
	return nil
}

func main() {
	checkPerson()
}

func checkPerson() {
	person := Person{
		Name: "Jane",
		Age:  18,
	}

	fmt.Printf("%v is an adult ? -- %t", person.Name, person.IsAdult())
}
