package main

import "github.com/Khalid-Nowaf/animal"

func main() {

	// Dog Part ---------------------
	dog := animal.Dog{}
	println(dog.String())
	dog.Bark()

	dog.SetName("Soka")
	dog.SetAge(3)

	println("After using the setters -> ")
	println(dog.String())
	println()

	// Cat Part -------------------------
	cat := animal.NewCat("moka", 1)
	println(cat.String())
	cat.Meow()

	println("After useing the dot operator ->")
	cat.Name = "Hoka"
	cat.Age = 2
	println(cat.String())

}
