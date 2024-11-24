package lemin

import (
	"Lemin-Project/tools"
	"Lemin-Project/variables"
	"log"
	"os"
	"strconv"
	"strings"
)

// Read file input
func ReadFile() []string {
	if len(os.Args) != 2 {
		log.Fatalln(variables.Errors["Args"])
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln(variables.Errors["Input"])
	}

	elements := strings.Split(string(file), "\n")

	return elements
}

// Find all(number of ants,rooms,tunnels)
func FindAll(Rooms *map[string]*variables.Room) {
	input := ReadFile()

	// Check the number of ants
	if tools.IsNumeric(input[0]) {
		var err error
		variables.NAnt, err = strconv.Atoi(input[0])
		if err != nil || variables.NAnt == 0 {
			log.Fatalln(variables.Errors["Ants"])
		}
	} else {
		log.Fatalln(variables.Errors["Ants"])
	}

	for i := 1; i < len(input); i++ {
		// Check repeated lines in file
		for j := 0; j < i; j++ {
			if input[j] == input[i] {
				log.Fatalln(variables.Errors["Repeat"])
			}
		}

		if strings.HasPrefix(input[i], "##") { // Check Start,End and Comments
			input[i] = strings.TrimPrefix(input[i], "##")
			if input[i] == "start" {
				str := strings.FieldsFunc(input[i+1], func(r rune) bool {
					return r == ' '
				})
				variables.Start = str[0]
			} else if input[i] == "end" {
				str := strings.FieldsFunc(input[i+1], func(r rune) bool {
					return r == ' '
				})
				variables.End = str[0]
			}
		} else if tools.IsRoom(input[i]) { // Check for Rooms
			name, Room := CompletRoom(input[i])
			(*Rooms)[name] = Room
		} else if tools.IsTunnel(input[i]) { // Check for tunnuls
			CompletRelation(input[i])
		}else if strings.HasPrefix(input[i], "#"){
			continue
		}  else {
			log.Fatalln(variables.Errors["Invalid!"])
		}
	}
}

// Add Room
func CompletRoom(line string) (string, *variables.Room) {
	Room := &variables.Room{}
	var Name string
	var err error
	parts := strings.Split(line, " ")

	Name = parts[0]
	Room.X, err = strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalln(variables.Errors["XError"])
	}
	Room.Y, err = strconv.Atoi(parts[2])
	if err != nil {
		log.Fatalln(variables.Errors["YError"])
	}
	Room.Visited = false

	return Name, Room
}

func CompletRelation(line string) {
	str := strings.FieldsFunc(line, func(r rune) bool {
		return r == '-'
	})

	if len(str) != 2 {
		log.Fatalln(variables.Errors["Tunnul"])
	}

	roomA, okA := variables.Rooms[str[0]]
	roomB, okB := variables.Rooms[str[1]]

	if !okA || !okB {
		log.Fatalln(variables.Errors["TunnulRoom"])
	}

	roomA.Relations = append(roomA.Relations, roomB)
}
