package main

func main() {

	V := st_JobSite{vacancies: []string{
		"mob_dev",
	}}
	bob := &Person{name: "bob"}
	V.addVacancies("Senior HTML Developer")
	V.Subscribe(bob)
	V.addVacancies("Java Junior Developer")
	V.removeVacancy("Senior HTML Developer")

}
