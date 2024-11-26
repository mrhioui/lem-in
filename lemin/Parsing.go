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
func ReadFile() string {
	if len(os.Args) != 2 {
		log.Fatalln(variables.Errors["Args"])
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln(variables.Errors["Input"])
	}
	return string(file)
}

// Find all(number of ants,rooms,tunnels)
func FindAll() {
	input := strings.Split(ReadFile(), "\n")

	// Check the number of ants
	if tools.IsNumeric(input[0]) {
		var err error
		variables.NAnt, err = strconv.Atoi(input[0])
		if err != nil || variables.NAnt == 0 {
			log.Fatalln(variables.Errors["Ants"])
		}
	} else {
		log.Fatalln(variables.Errors["AntsMissing"])
	}

	for i := 1; i < len(input); i++ {
		// Check repeated lines in file
		for j := 0; j < i; j++ {
			if input[j] == input[i] {
				log.Fatalln(variables.Errors["Repeat"] + input[j])
			}
		}
		if strings.HasPrefix(input[i], "##") { // Check Start,End and Comments
			input[i] = strings.TrimPrefix(input[i], "##")
			str := strings.Split(input[i+1], " ")
			if input[i] == "start" && tools.IsRoom(input[i+1]) {
				variables.Start = str[0]
			} else if input[i] == "end" && tools.IsRoom(input[i+1]) {
				variables.End = str[0]
			}else{
				log.Fatalln(variables.Errors["Start-End"])
			}
		} else if tools.IsRoom(input[i]) { // Check for Rooms
			CompletRoom(input[i])
		} else if tools.IsTunnel(input[i]) { // Check for tunnuls
			CompletRelation(input[i])
		} else if strings.HasPrefix(input[i], "#") {
			continue
		} else {
			log.Fatalln(variables.Errors["Invalid!"])
		}
	}
}

// Add Room
func CompletRoom(line string) {
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
	
	variables.Rooms[Name] = Room
}

func CompletRelation(line string) {
	str := strings.Split(line,"-")

	roomA, okA := variables.Rooms[str[0]]
	roomB, okB := variables.Rooms[str[1]]
	if !okA || !okB {
		log.Fatalln(variables.Errors["TunnulRoom"])
	}

	roomA.Relations = append(roomA.Relations, roomB)
	roomB.Relations = append(roomB.Relations, roomA)
}
