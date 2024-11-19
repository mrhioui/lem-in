package lemin

func dfsFindPaths(CurrentRoom Room, EndRoom Room, Path []Room, Paths *[][]Room) {
	// Visited = append(Visited, CurrentRoom)
	Path = append(Path, CurrentRoom)
	for _, relationRoom := range CurrentRoom.Relations {
		// fmt.Println(CurrentRoom.Name)
		for _, vis := range Path {
			if vis.Name == relationRoom.Name {
				continue
			}
		}
		if relationRoom.Name == EndRoom.Name {
			Path = append(Path, relationRoom)
			*Paths = append(*Paths, Path)
			Path = []Room{}
			// Visited = []Room{}
		}
		dfsFindPaths(relationRoom, EndRoom, Path, Paths)
	}

}

func FindPaths(Shema *Shema) {
	var Visited []Room
	dfsFindPaths(Shema.Start, Shema.End, Visited, &Shema.Paths)
}
