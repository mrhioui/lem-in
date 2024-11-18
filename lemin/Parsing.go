package lemin

import (
	"log"
	"os"
	"strconv"
	"strings"

	"Lemin-Project/tools"
)

// Read file input
func ReadFile() []string {
	if len(os.Args) != 2 {
		log.Fatalln("ERROR: There more argements not inportant OR no argement file")
		os.Exit(1)
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln("ERROR: unable to read file:", err)
		os.Exit(3)
	}

	elements := strings.Split(string(file), "\n")

	return elements
}

// find all(number of ants,rooms,tunnels)
func FindAll(input []string, Shema *Shema, Room *Room) {
	checkS, checkE := false, false
	for i := 0; i < len(input); i++ {
		for j := 0; j < i; j++ {
			if input[j] == input[i] {
				log.Fatalln("this line is repeated : " + input[i])
			}
		}
		switch {
		case i == 0 && tools.IsNumeric(input[i]):
			var err error
			Shema.NAnt, err = strconv.Atoi(input[0])
			if err != nil || Shema.NAnt == 0 {
				log.Fatalln("the number of antsis not correct or missing")
			}
		case tools.IsRoom(input[i]):
			Room := CompletRoom(input[i])
			for _, v := range Shema.Rooms {
				if v.Name == Room.Name || (v.X == Room.X && v.Y == Room.Y) {
					log.Fatalln("this room is repeated or cordenets : ", input[i])
				}
			}
			Shema.Rooms = append(Shema.Rooms, Room)
		case strings.HasPrefix(input[i], "##"):
			input[i] = strings.TrimPrefix(input[i], "##")
			if input[i] == "start" && !checkS {
				checkS = true
				i++
				if !tools.IsRoom(input[i]) {
					log.Fatalln("this is not a Room")
				}
				StartRoom := CompletRoom(input[i])
				Shema.Start = StartRoom
				Shema.Rooms = append(Shema.Rooms, StartRoom)
			} else if input[i] == "end" && !checkE {
				checkE = true
				i++
				if !tools.IsRoom(input[i]) {
					log.Fatalln("this is not a Room")
				}
				EndRoom := CompletRoom(input[i])
				Shema.End = EndRoom
				Shema.Rooms = append(Shema.Rooms, EndRoom)
			} else if (input[i] != "start") || (input[i] != "end") {
				continue
			} else {
				log.Fatalln("this line is repeated : ##" + input[i])
			}
		case tools.IsTunnel(input[i]):
			CompletTunnul(input[i], Shema.Rooms, &Shema.Start, &Shema.End)
		default:
			continue
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
func CompletTunnul(line string, Rooms []Room, Start *Room, End *Room) {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		log.Fatalln("Tunnel definition is invalid")
		return
	}

	var roomA *Room
	var roomB *Room

	// check if exist room
	for i := range Rooms {
		if Rooms[i].Name == parts[0] {
			roomA = &Rooms[i]
		}
		if Rooms[i].Name == parts[1] {
			roomB = &Rooms[i]
		}
	}

	// check if bad links
	if roomA == roomB {
		log.Fatalln("this tunnul have relation of one room: ", line)
	}
	if roomA != nil && roomB != nil {
		if roomA.Name == Start.Name {
			Start.Relations = append(Start.Relations, *roomB)
		} else if roomB.Name == Start.Name {
			Start.Relations = append(Start.Relations, *roomA)
		}
		if roomA.Name == End.Name {
			End.Relations = append(End.Relations, *roomB)
		} else if roomB.Name == End.Name {
			End.Relations = append(End.Relations, *roomA)
		}
		roomA.Relations = append(roomA.Relations, *roomB)
		roomB.Relations = append(roomB.Relations, *roomA)
	} else {
		log.Fatalln("One or both rooms not found for tunnel:", line)
	}
}
