package tools

import (
	"strconv"
	"strings"
)

func IsNumeric(s string) bool {
	test := false
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			test = true
		} else {
			test = false
			break
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
				return false
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
	} else {
		for v := range str {
			if !IsNumeric(strconv.Itoa(v)) {
				return false
			}
		}
	}
	return true
}
