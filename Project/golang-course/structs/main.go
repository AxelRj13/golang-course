package main

import "fmt"

// creating a struct, similar to object in JS
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
	// creating new instance of person struct
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// fmt.Printf("The name is %v %v", alex.firstName, alex.lastName)
	// fmt.Println()
	// fmt.Println()

	// other way to create new instance of the struct
	// var person1 person
	// fmt.Println(person1) // will only print out empty string because person1 struct's value has not been set to anything
	// person1.firstName = "Shay"
	// person1.lastName = "Cormac"
	// fmt.Printf("%+v", person1) // %+v means print out all different values from a struct
	// fmt.Println()
	// fmt.Println()

	// using embedded struct
	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Green",
	// 	contactInfo: contactInfo{
	// 		email:   "jim@gmail.com",
	// 		zipCode: 98217,
	// 	},
	// }
	var jim person
	jim.firstName = "Jim"
	jim.lastName = "Morris"
	jim.email = "jimorris@gmail.com"
	jim.zipCode = 54218

	// create pointer of jim structs so that when it passed as receiver, the real object / real structs will be updated accordingly.
	// If we do not use this, the updateName function will only update the copy of jim structs
	// jimPointer := &jim // give the memory address of the value this variable is pointing at
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy") // we can actually use the type of person directly without need to create new variable that point at the memory address
	jim.print()
	fmt.Println()

	name := "bill"
	namePointer := &name //turn name into memory address
	fmt.Println(namePointer)
	printPointer(namePointer)
}

// using structs as receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}

// update structs data with struct (pointer) as receiver
// *pointer is to give the value this memory address is pointing at
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func printPointer(namePointer *string) {
	fmt.Println(*namePointer)
}
