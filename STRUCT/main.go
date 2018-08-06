package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{"Alex", "Anderson"}

	// alex := person{ firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var james person
	// fmt.Println(james)
	// james.firstName = "Arijit"
	// james.lastName = "Roy"
	// // %+v prints out all the fieldnames and corresponding values in a struct
	// fmt.Printf("%+v", james);

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 4567732,
		},
	}
	// jimPointer := &jim;
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()

}
func (personPointer *person) updateName(newFirstName string) {
	(*personPointer).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
