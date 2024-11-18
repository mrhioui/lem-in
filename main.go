package main

import (
	"fmt"

	"Lemin-Project/lemin"
)

func main() {
	var Shema lemin.Shema
	var Room lemin.Room

	Variables := lemin.ReadFile()
	lemin.FindAll(Variables, &Shema, &Room)
	lemin.FindPaths(&Shema)
	// fmt.Printf("%v , NAnt\n", Shema.NAnt)
	fmt.Printf("%v , Start\n", Shema.Start)
	fmt.Printf("%v , End\n", Shema.End)
	// fmt.Printf("%v , Rooms\n", Shema.Rooms[0].Relations)
}
