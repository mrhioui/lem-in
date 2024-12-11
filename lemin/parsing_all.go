package lemin

import (
	"log"
	"os"
	"strconv"
	"strings"

	"Lemin-Project/tools"
	"Lemin-Project/variables"
)

// Parse all info (number of ants, rooms, tunnels)
func ParsingAll() string {
	var newInput []string
	input := readFile()

	input[0] = strings.TrimSpace(input[0])

	// Check the number of ants
	if tools.IsNumeric(input[0]) {
		var err error
		variables.NAnt, err = strconv.Atoi(input[0])
		if err != nil || variables.NAnt <= 0 {
			log.Fatalln(variables.Errors["Ants"])
		}
	} else {
		log.Fatalln(variables.Errors["Ants"])
	}
	newInput = append(newInput, input[0])

	for i := 1; i < len(input); i++ {
		if input[i] != "" {
			newInput = append(newInput, input[i])
		}
		input[i] = strings.TrimSpace(input[i])
		checkRepeat(input, i)
		conditions(&variables.Rooms, input, i)
	}

	if variables.Start == "" || variables.End == "" {
		log.Fatalln(variables.Errors["Invalid"])
	}

	return strings.Join(newInput, "\n")
}

// Read file input
func readFile() []string {
	if len(os.Args) != 2 {
		log.Fatalln(variables.Errors["Args"])
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln(variables.Errors["Input"])
	}

	elements := strings.TrimSpace(string(file))
	lines := strings.Split(elements, "\n")

	return lines
}

// Check repeated lines in file
func checkRepeat(input []string, i int) {
	for j := 0; j < i; j++ {
		if input[j] == input[i] && input[i] != "" {
			log.Fatalf(variables.Errors["Repeat"]+"%d", i)
		}
	}
}

func conditions(Rooms *map[string]*variables.Room, input []string, i int) {
	if strings.HasPrefix(input[i], "##") { // Check Start, End and Comments
		input[i] = strings.TrimPrefix(input[i], "##")
		if input[i] == "start" && i+1 < len(input) && len(input[i+1]) > 0 {
			str := strings.FieldsFunc(input[i+1], func(r rune) bool {
				return r == ' '
			})
			variables.Start = str[0]
		} else if input[i] == "end" && i+1 < len(input) && len(input[i+1]) > 0 {
			str := strings.FieldsFunc(input[i+1], func(r rune) bool {
				return r == ' '
			})
			variables.End = str[0]
		} else {
			log.Fatalln(variables.Errors["Invalid"] + ":  ##" + input[i])
		}
		input[i] = "##" + input[i] // For printing
	} else if tools.IsRoom(input[i]) { // Check for Rooms
		name, Room := tools.CompletRoom(input[i])
		(*Rooms)[name] = Room
	} else if tools.IsTunnel(input[i]) { // Check for Tunnuls
		tools.CompletRelation(input[i])
	} else if !strings.HasPrefix(input[i], "#") && !(input[i] == "") {
		log.Fatalln(variables.Errors["Invalid"])
	}
}
