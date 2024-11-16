package main

import (
	"Lemin-Project/lemin"
	"Lemin-Project/tools"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// find all(number of ants,rooms,tunnels)
func FindAll(input []string, Shema *lemin.Shema, Room *lemin.Room) {
	for i := 0; i < len(input); i++ {
		switch {
		case i == 0 && tools.IsNumeric(input[i]):
			var err error
			Shema.NAnt, err = strconv.Atoi(input[0])
			if err != nil || Shema.NAnt == 0 {
				log.Fatalln("the number of antsis not correct or missing")
				os.Exit(1)
			}
		case tools.IsRoom(input[i]):
			Room := CompletRoom(input[i])
			Shema.Rooms = append(Shema.Rooms, Room)
		case strings.HasPrefix(input[i], "#"):
			input[i] = strings.TrimPrefix(input[i], "#")
			if input[i] == "#start" {
				i++
				if !tools.IsRoom(input[i]) {
					log.Fatalln("this is not a Room")
					os.Exit(1)
				}
				StartRoom := CompletRoom(input[i])
				Shema.Start = StartRoom
				Shema.Rooms = append(Shema.Rooms, StartRoom)
			} else if input[i] == "#end" {
				i++
				if !tools.IsRoom(input[i]) {
					log.Fatalln("this is not a Room")
					os.Exit(1)
				}
				EndRoom := CompletRoom(input[i])
				Shema.End = EndRoom
				Shema.Rooms = append(Shema.Rooms, EndRoom)
			} else {
				continue
			}
		case tools.IsTunnel(input[i]):
			roomRelated := CompletTunnul(input[i], Shema.Rooms)
			Room.Relations = append(Room.Relations, roomRelated)
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
func CompletTunnul(line string, Rooms []lemin.Room) lemin.Room {
	var Room lemin.Room
	parths := strings.Split(line, "-")
	// fmt.Println(parths)
	for _, v := range Rooms {
		// fmt.Println("1 :", v)
		if v.Name == parths[0] {
			for _, a := range Rooms {
				// fmt.Println("2 :", a)
				if a.Name == parths[1] {
					// Room.Relations = append(Room.Relations, a)
					Room = a
				}
			}
		}
	}
	return Room
}

func main() {
	var Shema lemin.Shema
	var Room lemin.Room
	input := []string{"10", "##start", "start 1 6", "0 4 8", "o 6 8", "n 6 6", "e 8 4", "t 1 9", "E 5 9", "a 8 9", "m 8 6", "h 4 6", "A 5 2", "c 8 1", "k 11 2", "##end", "end 11 6", "start-t", "n-e", "a-m", "A-c", "0-o", "E-a", "k-end", "start-h", "o-n", "m-end", "t-E", "start-0", "h-A", "e-end", "c-k", "n-m", "h-n"}
	FindAll(input, &Shema, &Room)

	fmt.Printf("%v , NAnt\n", Shema.NAnt)
	fmt.Printf("%v , Start\n", Shema.Start)
	fmt.Printf("%v , End\n", Shema.End)
	fmt.Printf("%v , Rooms\n", Shema.Rooms)
	fmt.Printf("%v , relations\n", Shema.Start.Relations)

}
