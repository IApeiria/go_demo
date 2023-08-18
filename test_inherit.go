package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) sleep() {
	fmt.Println("sleep...")
}
func (a Animal) eat() {
	fmt.Println("eat...")
}

type Dog struct {
	Animal
	color string
}

func main() {
	dog := Dog{Animal{"puppy", 3},
		"blue",
	}
	dog.eat()
	dog.sleep()
	fmt.Println(dog.age)
	fmt.Println(dog.color)
}
