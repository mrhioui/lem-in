package lemin

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Read file input
func ReadFile() []string {
	if len(os.Args) != 2 {
		log.Println("ERROR: There more argements not inportant OR no argement file")
		os.Exit(1)
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
	} else {
		fmt.Println("ERROR: unable to read file:", err)
		os.Exit(3)
	}

	elements := strings.Split(string(file), " ")

	return elements
}

// find all(number of ants,rooms,tunnels)
func Find() {
	input := ReadFile()

	for _, v := range input {
		if v == "##start" {
			
		}
	}
}
