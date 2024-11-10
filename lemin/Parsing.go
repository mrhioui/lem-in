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
func FindAll(input []string, Shema *Shema) {
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
			Tunnul := CompletTunnul(input[i], Shema.Rooms)
			Shema.Tunnuls = append(Shema.Tunnuls, Tunnul)
		default:
			continue
		}
	}
}

// Add Name and Point of Room
func CompletRoom(line string) Room {
	var Room Room
	var err error
	parts := strings.FieldsFunc(line, func(r rune) bool {
		return r == ' '
	})

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

// // Add Start and End of Tunnul
func CompletTunnul(line string, Rooms []Room) Tunnel {
	var Tunnul Tunnel
	parths := strings.FieldsFunc(line, func(r rune) bool {
		return r == '-'
	})

	for _, v := range Rooms {
		if v.Name == parths[0] {
			Tunnul.From = v
		} else if v.Name == parths[1] {
			Tunnul.To = v
		}
	}

	if Tunnul.From.Name == "" || Tunnul.To.Name == "" {
		log.Fatalln("There are a rome not exist here !!")
		os.Exit(1)
	}
	return Tunnul
}
