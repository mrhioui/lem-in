package tools

import (
	"log"
	"strconv"
	"strings"

	"Lemin-Project/variables"
)

func IsNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] <= '0' && s[i] >= '9' {
			return false
		}
	}
	return true
}

func IsRoom(s string) bool {
	str := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' '
	})
	if len(str) != 3 {
		return false
	} else {
		if str[0] == "L" || strings.HasPrefix(str[0], "#") {
			log.Fatalln(variables.Errors["Room"] + str[0])
		}
	}
	return true
}

func IsTunnel(s string) bool {
	str := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-'
	})
	if len(str) != 2 {
		if strings.Contains(s, "-") {
			log.Fatalln(variables.Errors["Tunnul"])
		}
		return false
	}
	return true
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
		log.Fatalf(variables.Errors["XError"] + line)
	}
	Room.Y, err = strconv.Atoi(parts[2])
	if err != nil {
		log.Fatalln(variables.Errors["YError"] + line)
	}
	Room.Visited = false

	checkCordinents(Name, Room.X, Room.Y)

	return Name, Room
}

func checkCordinents(name string, x, y int) {
	for n, v := range variables.Rooms {
		if v.X == x && v.Y == y {
			log.Fatalf("%s (%s %d %d) (%s %d %d)", variables.Errors["XYError"], n, v.X, v.Y, name, x, y)
		}
	}
}

// Add Relation
func CompletRelation(line string) {
	str := strings.Split(line, "-")

	roomA, okA := variables.Rooms[str[0]]
	roomB, okB := variables.Rooms[str[1]]

	if !okA || !okB {
		log.Fatalln(variables.Errors["TunnulRoom"], line)
	}

	roomA.Relations = append(roomA.Relations, roomB)
	roomB.Relations = append(roomB.Relations, roomA)
}
