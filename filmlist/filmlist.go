package film

import "fmt"

type Film struct {
	Name     string
	Name_eng string
	Date     string
	Vote     string
}

func New(vote, name, name_eng, date string) *Film {

	f := new(Film)

	f.Name = name
	f.Name_eng = name_eng
	f.Date = date
	f.Vote = vote

	return f
}

func (this *Film) String() string {
	return fmt.Sprintf("Название на русском:\t\t%s\nНазвание на английском:\t\t%s\nДата оценки:\t\t%s\nОценка:\t\t%s",
		this.Name, this.Name_eng, this.Date, this.Vote)
}
