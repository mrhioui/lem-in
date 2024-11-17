package main

import (
	"Lemin-Project/lemin"
	"Lemin-Project/tools"
	"log"
	"strconv"
	"strings"
)

// find all(number of ants,rooms,tunnels)
func FindAll(input []string, Shema *lemin.Shema, Room *lemin.Room) {
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
			} else if input[i] != "start" || input[i] != "end" {
				continue
			} else {
				log.Fatalln("this line is repeated : ##" + input[i])
			}
		case tools.IsTunnel(input[i]):
			CompletTunnul(input[i], Shema.Rooms)
		default:
			continue
		}
	}
}

// Add Name and Point of Room
func CompletRoom(line string) lemin.Room {
	var Room lemin.Room
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
func CompletTunnul(line string, Rooms []lemin.Room) {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		log.Fatalln("Tunnel definition is invalid")
		return
	}

	var roomA *lemin.Room
	var roomB *lemin.Room

	for i := range Rooms {
		if Rooms[i].Name == parts[0] {
			roomA = &Rooms[i]
		}
		if Rooms[i].Name == parts[1] {
			roomB = &Rooms[i]
		}
	}

	if roomA == roomB {
		log.Fatalln("this tunnul have relation of one room: ",line)
	}
	if roomA != nil && roomB != nil {
		roomA.Relations = append(roomA.Relations, *roomB)
		roomB.Relations = append(roomB.Relations, *roomA)
	} else {
		log.Fatalln("One or both rooms not found for tunnel:", line)
	}
}

func main() {
	var Shema lemin.Shema
	var Room lemin.Room
	input := []string{"4", "##start", "0 0 3", "2 2 5", "3 4 0", "##end", "1 8 3", "0-2", "2-3", "3-1"}
	FindAll(input, &Shema, &Room)

	// fmt.Printf("%v , NAnt\n", Shema.NAnt)
	// fmt.Printf("%v , Start\n", Shema.Start)
	// fmt.Printf("%v , End\n", Shema.End)
	// fmt.Println(Shema.Rooms)
	// for _, v := range Shema.Rooms {
	// 	fmt.Println(v.Name, v.Relations)
	// }
}
