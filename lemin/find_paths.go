package lemin

import (
	"log"
	"sort"

	"Lemin-Project/variables"
)

// DFS explore a path as deeply as possible before backtracking
func FindPaths(currentRoom *variables.Room, Qpath []string) {
	defer func() {
		currentRoom.Visited = false
	}()
	currentRoom.Visited = true
	Qpath = append(Qpath, getRoomName(currentRoom))

	for _, childRoom := range currentRoom.Relations {
		if !childRoom.Visited {
			if childRoom == variables.Rooms[variables.End] {
				Qpath = append(Qpath, variables.End)
				copyPath := append([]string{}, Qpath...)
				variables.Paths = append(variables.Paths, copyPath)
				return
			} else {
				FindPaths(childRoom, Qpath)
			}
		}
	}
}

func getRoomName(currentRoom *variables.Room) string {
	for rName, room := range variables.Rooms {
		if room == currentRoom {
			return rName
		}
	}
	return ""
}

func SortPaths(paths [][]string) {
	pathsCheck()
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
}

func pathsCheck() {
	// No path was found
	if len(variables.Paths) == 0 {
		log.Fatalln(variables.Errors["Invalid"])
	}
	// An incomplete path
	for i := range variables.Paths {
		if len(variables.Paths[i]) < 2 {
			log.Fatalln(variables.Errors["Invalid"])
		}
	}
}
