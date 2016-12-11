package main

import e "github.com/Khalid-Nowaf/embedded"

func main() {
	a := e.Derived{e.Base{"Base-a", 10, 32.2}, 2.5}
	a.Display()
	a.Base.Display()
	a.Xyz()
}
