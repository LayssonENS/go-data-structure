package main

import "fmt"

type Person struct {
	Age    int
	Height float32
}

func main() {

	newPerson := Person{}
	newPerson.Height = 1.85
	newPerson.Age = 23

	fmt.Println("Idade =", newPerson.Age)
	fmt.Println("Altura =", newPerson.Height)

}
