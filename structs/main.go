package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email: "jim@ex.com",
			zipCode: 79000,
		},
	}

	var jill person

	jill.firstName = "Jill"
	jill.lastName = "Party"
	jill.contact.email = "jill@ex.com"
	jill.contact.zipCode = 87654

	fmt.Println(jill)
	fmt.Println(jim)

	jim.updateName("Jimmy")
	fmt.Println(jim)

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
