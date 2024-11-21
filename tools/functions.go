package tools

import (
	"log"
	"strings"
)

func IsNumeric(s string) bool {
	test := false
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			test = true
		} else {
			test = false
		}
	}
	return test
}

func IsRoom(s string) bool {
	str := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' '
	})
	if len(str) != 3 {
		return false
	} else {
		for i, v := range str {
			if i == 0 && v == "L" {
				HandleErrors(3, s)
			}
		}
	}
	return true
}

func IsTunnel(s string) bool {
	str := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-'
	})
	if len(str) != 2 {
		return false
	}
	return true
}

func HandleErrors(n int, s string) {
	switch n {
	case 0:
		log.Fatalln("ERROR: There more argements not inportant OR no argement file")
	case 1:
		log.Fatalln("the number of ants is not correct or missing")
	case 2:
		log.Fatalln("this line is repeated : " + s)
	case 3:
		log.Fatalln("this is not a room : " + s)
	case 4:
		log.Fatalln("this room is have repeated name or cordenets : " + s)
	case 5:
		log.Fatalln("ERROR: unable to read file: no such file or directory")
	}
}
