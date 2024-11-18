package main

import (
	"Lemin-Project/lemin"
)

func main() {
	var Shema lemin.Shema
	var Room lemin.Room

	Variables := lemin.ReadFile()
	lemin.FindAll(Variables, &Shema, &Room)
	lemin.FindPaths(&Shema)

	// fmt.Printf("%v , NAnt\n", Shema.NAnt)
	// fmt.Printf("%v , Start\n", Shema.Start)
	// fmt.Printf("%v , End\n", Shema.End)
	// fmt.Printf("%v , Rooms\n", Shema.Rooms[0])
	// fmt.Printf("%v , Rooms\n", Shema.Rooms[1])
	// fmt.Printf("%v , Rooms\n", Shema.Rooms[2])
	// fmt.Printf("%v , Rooms\n", Shema.Rooms[3])
	// fmt.Println("paths :", Shema.Paths)
	// for _, room := range Shema.Rooms {
	// 	fmt.Printf("Room: %v, Relations: ", room.Name)
	// 	for _, rel := range room.Relations {
	// 		fmt.Printf("%v ", rel.Name)
	// 	}
	// 	fmt.Println()
	// }

}
