package lemin

import (
	"Lemin-Project/tools"
	"log"
	"os"
	"strconv"
	"strings"
)

// Read file input
func ReadFile() []string {
	if len(os.Args) != 2 {
		tools.HandleErrors(0, "")
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		tools.HandleErrors(5, "")
	}

	elements := strings.Split(string(file), "\n")

	return elements
}

// Find all(number of ants,rooms,tunnels)
func FindAll(Shema *Shema, room *Room) {
	input := ReadFile()
	checkS, checkE := false, false
	var StartRoom, EndRoom string
	// var err int
	for i := 0; i < len(input); i++ {
		// Check repeated lines in file
		for j := 0; j < i; j++ {
			if input[j] == input[i] {
				tools.HandleErrors(2, input[i])
			}
		}

		// Check the number of ants
		if i == 0 {
			if tools.IsNumeric(input[i]) {
				var err error
				Shema.NAnt, err = strconv.Atoi(input[0])
				if err != nil || Shema.NAnt == 0 {
					tools.HandleErrors(1, "")
				}
			} else {
				tools.HandleErrors(1, "")
			}
		}

		// Check Start,End and Comments
		if strings.HasPrefix(input[i], "##") {
			input[i] = strings.TrimPrefix(input[i], "##")
			if input[i] == "start" && !checkS {
				checkS = true
				str := strings.FieldsFunc(input[i+1], func(r rune) bool {
					return r == ' '
				})
				StartRoom = str[0]
			} else if input[i] == "end" && !checkE {
				checkE = true
				str := strings.FieldsFunc(input[i+1], func(r rune) bool {
					return r == ' '
				})
				EndRoom = str[0]
			}
		}

		// Check for Rooms
		if tools.IsRoom(input[i]) {
			room := CompletRoom(input[i])
			for _, v := range Rooms {
				if v.Name == room.Name || (v.X == room.X && v.Y == room.Y) {
					tools.HandleErrors(4, input[i])
				}
			}
			Rooms = append(Rooms, room)
		}

		// Check for Tunnuls
		if tools.IsTunnel(input[i]) {
			// CompletTunnul(input[i], Rooms)
		}
	}

	// complet Start and End Room
	for _, room := range Rooms {
		if room.Name == StartRoom {
			Shema.Start = room
		}
		if room.Name == EndRoom {
			Shema.End = room
		}
	}
}

// Add Name and Point of Room
func CompletRoom(line string) Room {
	var Room Room
	var err error
	parts := strings.Split(line, " ")

	Room.Name = parts[0]
	Room.X, err = strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalln("the X of room not a number")
	}
	Room.Y, err = strconv.Atoi(parts[2])
	if err != nil {
		log.Fatalln("the Y of room not a number")
	}

	return Room
}

// Add Start and End of Tunnul
// func CompletTunnul(line string, Rooms []Room) {
// 	parts := strings.Split(line, "-")
// 	if len(parts) != 2 {
// 		log.Fatalln("Tunnel definition is invalid")
// 	}

// 	// var roomA *Room
// 	// var roomB *Room

// 	// check if exist room
// 	for i := range Rooms {
// 		if Rooms[i].Name == parts[0] {
// 			roomA = &Rooms[i]
// 		}
// 		if Rooms[i].Name == parts[1] {
// 			roomB = &Rooms[i]
// 		}
// 	}

// 	// check if bad links
// 	if roomA == roomB {
// 		log.Fatalln("this tunnul have relation of one room: ", line)
// 	}
// 	if roomA != nil && roomB != nil {
// 		roomA.Relations = append(roomA.Relations, roomB)
// 		roomB.Relations = append(roomB.Relations, roomA)
// 	} else {
// 		log.Fatalln("One or both rooms not found for tunnel:", line)
// 	}
// }
