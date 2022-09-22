package main

type st_JobSite struct {
	subcribers []Observer
	vacancies  []string
}

func JobSite() st_JobSite {
	return st_JobSite{
		subcribers: make([]Observer, 0),
		vacancies:  make([]string, 0),
	}
}
func (V st_JobSite) addVacancies(vacanci string) {
	V.vacancies = append(V.vacancies, vacanci)
	V.SendAll()
}

func (V st_JobSite) removeVacancy(vacanci string) {
	var i = 0
	var n = len(V.subcribers)
	for j, pt := range V.vacancies {
		if pt == vacanci {
			i = j
			break
		}
	}
	V.vacancies[i] = V.vacancies[n-1]
	V.vacancies[n-1] = ""
	V.vacancies = V.vacancies[:n-1]
	V.SendAll()
}
func (V st_JobSite) Subscribe(observer Observer) {
	V.subcribers = append(V.subcribers, observer)
}
func (V st_JobSite) Unsubscribe(observer Observer) {
	var i = 0
	var n = len(V.subcribers)
	for j, uns := range V.subcribers {
		if uns == observer {
			i = j
			break
		}
	}
	V.subcribers[i] = V.subcribers[n-1]
	V.subcribers[n-1] = nil
	V.subcribers = V.subcribers[:n-1]
}
func (V st_JobSite) SendAll() {
	for _, sen := range V.subcribers {
		sen.handleEvent(V.vacancies)
	}

}
