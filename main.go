package main

import (
	"Lemin-Project/lemin"
	"fmt"
)

func main() {
	//parse input, validate, find paths, simulate movements
	var Shema lemin.Shema
	Variables := lemin.ReadFile()
	lemin.FindAll(Variables, &Shema)

	fmt.Printf("%v , NAnt\n", Shema.NAnt)
	fmt.Printf("%v , Start\n", Shema.Start)
	fmt.Printf("%v , End\n", Shema.End)
	fmt.Printf("%v , Rooms\n", Shema.Rooms)
	for i := 0; i < len(Shema.Tunnuls); i++ {
		fmt.Printf("%v\n", Shema.Tunnuls[i])
	}
}
