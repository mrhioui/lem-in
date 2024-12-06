package lemin

import (
	"log"
	"sort"

	"Lemin-Project/tools"
	"Lemin-Project/variables"
)

var Paths [][]string

// DFS exploring a path as deeply as possible before backtracking
func ExtractPaths(currentRoom *variables.Room, Qpath []string) {
	defer func() {
		currentRoom.Visited = false
	}()
	currentRoom.Visited = true
	Qpath = append(Qpath, tools.GetRoomName(currentRoom))

	for _, childRoom := range currentRoom.Relations {
		if !childRoom.Visited {
			if childRoom == variables.Rooms[variables.End] {
				Qpath = append(Qpath, variables.End)
				copyPath := append([]string{}, Qpath...)
				Paths = append(Paths, copyPath)
				return
			} else {
				ExtractPaths(childRoom, Qpath)
			}
		}
	}
}

func SortPaths(paths [][]string) {
	pathsCheck()
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
}

func pathsCheck() {
	// No path was found
	if len(Paths) == 0 {
		log.Fatalln(variables.Errors["Invalid"])
	}
	// An incomplete path
	for i := range Paths {
		if len(Paths[i]) < 2 {
			log.Fatalln(variables.Errors["Invalid"])
		}
	}
}
