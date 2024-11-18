package lemin

import "fmt"

func FindPaths(Shema *Shema)  {
	// var Paths [][]Room
	// var Path Room

	for _, v := range Shema.Start.Relations {
		fmt.Println("room :", v.Name)
		// Path = v
	}
}
