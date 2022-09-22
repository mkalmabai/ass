package main

import "fmt"

type Person struct {
	name string
}

func (x Person) handleEvent(vacancies []string) {
	fmt.Println("Hi dear ", x.name)
	fmt.Println("Vacancies updated: ")
	for _, vacancy := range vacancies {
		fmt.Println(vacancy)
	}

}
