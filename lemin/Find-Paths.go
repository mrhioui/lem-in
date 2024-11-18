package lemin

func dfsFindPaths(CurrentRoom Room, EndRoom Room, Visited map[string]bool, Path []Room, Paths *[][]Room) {
	Visited[CurrentRoom.Name] = true
	Path = append(Path, CurrentRoom)

	for _, relationRoom := range CurrentRoom.Relations {
		if Visited[relationRoom.Name] {
			continue
		} else if relationRoom.Name != EndRoom.Name {
			dfsFindPaths(relationRoom, EndRoom, Visited, Path, Paths)
		} else {
			*Paths = append(*Paths, Path)
		}
	}
}

// FindPaths finds all paths from Start to End in a Shema
func FindPaths(Shema *Shema) {
	Visited := make(map[string]bool)
	dfsFindPaths(Shema.Start, Shema.End, Visited, []Room{}, &Shema.Paths)
}
