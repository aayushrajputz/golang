package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

type contact struct {
	phone int
	email string
}
type address struct {
	HouseNo int
	city    string
	state   string
	country string
}

type employee struct {
	person_details  person
	contact_details contact
	address_details address
}

func main() {
	var aayush person
	fmt.Println(aayush)
	aayush.firstname = "aayush"
	aayush.lastname = "rajput"
	aayush.age = 20

	fmt.Println(aayush)

	person1 := person{
		firstname: "akash",
		lastname:  "rajput",
		age:       20,
	}
	fmt.Println(person1, "this is person1")

	//new keyword
	person2 := new(person)
	person2.firstname = "punit"
	person2.lastname = "rajput"
	person2.age = 20

	fmt.Println(person2, "this is person2 data")

	employee1 := employee{
		person_details: person{
			firstname: "aayush",
			lastname:  "rajput",
			age:       20,
		},
		contact_details: contact{
			phone: 1234567890,
			email: "aayush@gmail.com",
		},
		address_details: address{
			HouseNo: 123,
			city:    "delhi",
			state:   "delhi",
			country: "india",
		},
	}

	fmt.Println(employee1.contact_details)

}
