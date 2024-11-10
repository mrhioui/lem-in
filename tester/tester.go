package main

import (
	"fmt"
	"Lemin-Project/lemin"
	"Lemin-Project/tools"
	"log"
	"os"
	"strconv"
	"strings"
)

// find all(number of ants,rooms,tunnels)
func FindAll(input []string, Shema *lemin.Shema) {
	for i := 0; i < len(input); i++ {
		switch {
		case i == 0 && tools.IsNumeric(input[0]):
			var err error
			Shema.NAnt, err = strconv.Atoi(input[0])
			if err != nil {
				log.Fatalln("the number of antsis not correct or missig")
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
			log.Fatalln("not yet thinking about message")
			os.Exit(1)
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

// // Add Start and End of Tunnul
func CompletTunnul(line string, Rooms []lemin.Room) lemin.Tunnel {
	var Tunnul lemin.Tunnel
	parths := strings.Split(line, "-")

	for _, v := range Rooms {
		if v.Name == parths[0] {
			Tunnul.From = v
		} else if v.Name == parths[1] {
			Tunnul.To = v
		}
	}

	if Tunnul.From.Name == "" || Tunnul.To.Name == "" {
		log.Fatalln("There a rome not exist here !!")
		os.Exit(1)
	}
	return Tunnul
}

func main() {
	var Shema lemin.Shema
	input := []string{
		"2", "##start", "1 23 3", "2 16 7", "#comment", "3 16 3", "4 16 5",
		"5 9 3", "6 1 5", "7 4 8", "##end", "0 9 5", "0-4", "0-6", "1-3",
		"4-3", "5-2", "3-5", "#another comment", "4-2", "2-1", "7-6",
		"7-2", "7-4", "6-5"}

	FindAll(input, &Shema)

	fmt.Printf("%v , NAnt\n", Shema.NAnt)
	fmt.Printf("%v , Start\n", Shema.Start)
	fmt.Printf("%v , End\n", Shema.End)
	fmt.Printf("%v , Rooms\n", Shema.Rooms)
	for i := 0;i<len(Shema.Tunnuls);i++{
		fmt.Printf("%v\n", Shema.Tunnuls[i])
	}
}
