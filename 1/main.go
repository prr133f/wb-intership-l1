package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) String() string {
	return fmt.Sprintf("%s (%d years old)", h.Name, h.Age)
}

type Action struct {
	Human
}

func main() {
	bob := Human{
		Name: "Bob",
		Age:  42,
	}
	fmt.Println(bob.String())

	a := Action{Human: bob}
	fmt.Println(a.String())
}
