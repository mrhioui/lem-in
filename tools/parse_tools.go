package tools

import (
	"log"
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
