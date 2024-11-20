package main

import (
	"Lemin-Project/lemin"
	"fmt"
)

func main() {
	var Shema lemin.Shema
	var Room lemin.Room

	lemin.FindAll(&Shema, &Room)
	// lemin.FindPaths(&Shema)

	fmt.Println(Shema.NAnt)
	fmt.Println(Shema.Start)
	fmt.Println(Shema.End)
	// fmt.Println(Shema.Paths)
}
