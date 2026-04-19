package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	//contact   contactInfo
	contactInfo
}

type contactInfo struct {
	phone   int
	email   string
	pincode int
}

func main() {
	// var ajain person
	// ajain.firstname = "Ajain"
	// ajain.lastname = "Krishnan"
	//	ajain := person{"Ajain", "Krishnan"}
	ajain := person{
		firstname: "Ajain",
		lastname:  "Krishnan",
		// contact: contactInfo{
		// 	phone:   1234567890,
		// 	email:   "ajain@example.com",
		// 	pincode: 123456,
		// },
		contactInfo: contactInfo{
			phone:   1234567890,
			email:   "ajain@example.com",
			pincode: 123456,
		},
	}
	ajain.print()

	ajain.updateName("Amalda")
	//pointerToAjain := &ajain
	//pointerToAjain.updateName("Amalda")
	ajain.print()
}

//	func (p person) updateName(newFirstName string) {
//		p.firstname = newFirstName
//		p.print()
//	}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName
	pointerToPerson.print()
}

func (p person) print() {
	fmt.Printf("Welcome %v %v to the world of Go!\n", p.firstname, p.lastname)
	//fmt.Println(p)
}
