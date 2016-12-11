package main

import (
	"fmt"

	e "github.com/Khalid-Nowaf/embedded"
)

func main() {

	var x e.Derived
	fmt.Printf("%T\n", x.A)
	fmt.Printf("%T\n", x.Base.A)
	fmt.Printf("%T\n", x.C)
	fmt.Printf("%T\n", x.Base.C)

}
