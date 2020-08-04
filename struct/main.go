package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//1st way
	//r := person{firstName: "rohit", lastName: "k"}
	//fmt.Println(r)

	//2nd way
	//var ro person
	//ro.firstName = "R"
	//ro.lastName = "k"

	//fmt.Println(ro)
	//fmt.Printf("%+v", ro)

	//3rd Way
	roh := person{
		firstName: "r",
		lastName:  "k",
		contactInfo: contactInfo{
			email:   "blah@gmail.com",
			zipcode: 122334,
		},
	}
	rohPointer := &roh
	rohPointer.updateName("rohit")
	// or
	roh.updateName("RohitK")

	roh.print()
}

func (personToPointer *person) updateName(newFirstName string) {
	(*personToPointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
