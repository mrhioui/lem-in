package lemin

import (
	"sort"

	"Lemin-Project/variables"
)

var Paths [][]string

func getRoomName(currentRoom *variables.Room) string {
	for rName, room := range variables.Rooms {
		if room == currentRoom {
			return rName
		}
	}
	return ""
}

func ExtractPaths(currentRoom *variables.Room, Qpath []string) {
	defer func() {
		currentRoom.Visited = false
	}()
	currentRoom.Visited = true
	Qpath = append(Qpath, getRoomName(currentRoom))
	if currentRoom == variables.Rooms[variables.End] {
		Paths = append(Paths, Qpath)
	} else {
		for _, childRoom := range currentRoom.Relations {
			if !childRoom.Visited {
				if childRoom == variables.Rooms[variables.End] {
					Paths = append(Paths, append(Qpath, variables.End))
					return
				} else {
					ExtractPaths(childRoom, Qpath)
				}
			}
		}
	}
}

func SortPaths() {
	sort.Slice(Paths, func(i, j int) bool {
		return len(Paths[i]) < len(Paths[j])
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
