package main

type Person1 struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person1, error) {
	if name == "" {
		return nil, fmt.Errof(name)
	}
}
