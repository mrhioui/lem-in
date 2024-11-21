package main

import (
	"Lemin-Project/lemin"
	"fmt"
)

func main() {
	var Shema lemin.Shema
	var Room lemin.Room
	
	Room= lemin.Room{
		X: 56, 
		Y: 76,
		Visited: false,
	}
	
	Rooms["0"] = Room
	lemin.FindAll(&Shema, &Room)
	// lemin.FindPaths(&Shema)

	fmt.Println(Shema.NAnt)
	fmt.Println(Shema.Start)
	fmt.Println(Shema.End)
	// fmt.Println(Shema.Paths)
}
