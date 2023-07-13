package person

import "fmt"

// Structs

type Person struct {
	FirstName string
	LastName  string
	email     string
}

// Exposed method
func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func NewPerson(fname, lname, email string) *Person {
	return &Person{
		FirstName: fname,
		LastName:  lname,
		email:     email,
	}
}
