package test

import "fmt"

type Animal interface {
	Sound()
	BreedOf()
}

type Dog struct {
	Name  string
	Age   int
	Breed string
}

func NewAnimal(name string, age int, breed string) Animal {
	return &Dog{
		Name:  name,
		Age:   age,
		Breed: breed,
	}
}

func (d *Dog) Sound() {
	fmt.Printf("%s says: Woof Woof!\n", d.Name)
}

func (d *Dog) BreedOf(){
	fmt.Printf("%s about: This breed is\n", d.Breed)
}