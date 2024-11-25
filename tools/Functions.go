package tools

import (
	"Lemin-Project/variables"
	"log"
	"strings"
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
			log.Fatalln(variables.Errors["Room"] + s)
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

func AreRoomsEqual(r1, r2 *variables.Room) bool {
	if r1 == nil || r2 == nil {
		return r1 == r2
	}
	
	if r1.X != r2.X || r1.Y != r2.Y || r1.Visited != r2.Visited || len(r1.Relations) != len(r2.Relations) {
		return false
	}

	for i := range r1.Relations {
		if !AreRoomsEqual(r1.Relations[i], r2.Relations[i]) {
			return false
		}
	}

	return true
}
