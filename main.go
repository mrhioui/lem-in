package main

import (
	"Lemin-Project/lemin"
	"Lemin-Project/variables"
	"fmt"
)

func main() {
	lemin.FindAll(&variables.Rooms)

	for key, value := range variables.Rooms {
		fmt.Printf("%v,  %p,  %v\n", key, value, value)
	}
}
