package main

import (
	"fmt"

	e "github.com/Khalid-Nowaf/embedded"
)

func main() {

	aRect := e.Rectangle{
		e.NamedObj{"name1"},
		e.Shape{e.NamedObj{"name2"}, 0, true},
		e.Point{0, 0},
		20, 2.5}

	fmt.Println(aRect.Name)
	fmt.Println(aRect.Shape)
	fmt.Println(aRect.Shape.Name)

}
