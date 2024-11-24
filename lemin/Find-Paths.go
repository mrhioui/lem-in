package lemin

import (
	"Lemin-Project/variables"
)

func FindAllPaths() {
	var path []variables.Room
	path = append(path, *variables.Rooms[variables.Start])

	if variables.Start == variables.End {
		variables.Paths = append(variables.Paths, path)
	} else {
		for _, v := range variables.Paths[len(variables.Paths)-1] {
			if (!path.co)
		}
	}
}
