package main

import (
	"fmt"

	"github.com/Khalid-Nowaf/animal"
)

func main() {
	animals := []animal.Animal{animal.NewDog("soka", 3), animal.NewCat("moka", 2), animal.Llama{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak()) // method dispatch via jmp-table
	}
}
