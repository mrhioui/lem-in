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
	if appendPath(currentRoom, Qpath) {
		return
	}
	for _, childRoom := range currentRoom.Relations {
		if !childRoom.Visited {
			if appendPath(childRoom, append(Qpath, variables.End)) {
				return
			}
			ExtractPaths(childRoom, Qpath)
		}
	}
}

func appendPath(currentRoom *variables.Room, Qpath []string) bool {
	if currentRoom == variables.Rooms[variables.End] {
		Paths = append(Paths, Qpath)
		return true
	}
	return false
}

func pathsCheck() {
	if len(Paths) == 0 {
		log.Fatalln(variables.Errors["Invalid"])
	}
	for i := range Paths {
		if len(Paths[i]) < 2 {
			log.Fatalln(variables.Errors["Invalid"])
		}
	}
}

func SortPaths(paths [][]string) {
	pathsCheck()
	sort.Slice(paths, func(i, j int) bool {
		return len(paths[i]) < len(paths[j])
	})
}

// -------------------------
// using room pointer as indicator
// -------------------------
// var Paths [][]*variables.Room

// func ExtractPaths(currentRoom *variables.Room, Qpath []*variables.Room) {
// 	defer func() {
// 		currentRoom.Visited = false
// 	}()
// 	currentRoom.Visited = true
// 	Qpath = append(Qpath, currentRoom)
// 	if currentRoom == variables.Rooms[variables.End] {
// 		Paths = append(Paths, Qpath)
// 	} else {
// 		for _, childRoom := range currentRoom.Relations {
// 			if !childRoom.Visited {
// 				ExtractPaths(childRoom, Qpath)
// 			}
// 		}
// 	}
// }
