package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (per Person) eat() {
	fmt.Println("eat...")
}
func (per Person) sleep() {

}
func main() {
	per := Person{
		name: "tome",
		age:  20,
	}
	per.eat()
}
