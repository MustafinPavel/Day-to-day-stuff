package main

import "fmt"

func main() {
	manOne := People{
		name:    "Michael",
		age:     23,
		country: "RF",
	}
	fmt.Println(manOne)
	manOne.register()
}

type People struct {
	name    string
	age     int
	country string
}

func (p People) register()int {
	fmt.Printf("%v was registered\n", p.name)
	return 0
}
