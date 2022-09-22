package main

type Observer interface {
	handleEvent(vacancies []string)
}
