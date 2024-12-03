package main

import (
	"fmt"
	"log"
	"strings"

	"Lemin-Project/lemin"
	"Lemin-Project/tools"
	"Lemin-Project/variables"
)

func PrintInfos() {
	if len(variables.Rooms) == 0 {
		log.Fatalln(variables.Errors["Invalid"])
	}

	reverseRooms := make(map[*variables.Room]string)
	for key, value := range variables.Rooms {
		reverseRooms[value] = key
	}

	for key, value := range variables.Rooms {
		fmt.Printf("* Name: %v\nAddress: %p\n", key, value)
		fmt.Printf("Coordinates: (X: %v, Y: %v)\n", value.X, value.Y)
		fmt.Printf("Visited: %v\n", value.Visited)
		fmt.Println("Relations:")
		for _, rel := range value.Relations {
			if name, found := reverseRooms[rel]; found {
				fmt.Printf("\t- %v (%p)\n", name, rel)
			} else {
				fmt.Printf("\t- Invalid relation (%p)\n", rel)
			}
		}
		fmt.Println(strings.Repeat("-", 40))
	}
}

func main() {
	// Parsing all info from file
	input := lemin.FindAll(&variables.Rooms)
	PrintInfos()
	fmt.Println("\n", strings.Repeat("-", 120))
	lemin.ExtractPaths(variables.Rooms[variables.Start], nil)
	// lemin.ExtractPaths(variables.Rooms[variables.Start],[]*variables.Room{})
	lemin.SortPaths()
	for i, path := range lemin.Paths {
		fmt.Printf("path n%d: %v >> capcity: %d\n", i+1, path, len(path)-2)
	}
	fmt.Println("\n", strings.Repeat("-", 120))
	combined := ValidPathsCombs()
	for i, com := range combined {
		fmt.Printf("combination n%d: %v >> number of combined paths %d\n", i+1, com, len(com))
	}
	fmt.Println("\n", strings.Repeat("-", 120))

	combined = tools.RemoveMatching(combined)

	fmt.Println(input)
	fmt.Println()

	combinations := lemin.BestPath(variables.NAnt, combined)
	lemin.AntMove(combinations, variables.NAnt)
}

// Combinations of non-intersecting paths
func ValidPathsCombs() (validCombs [][][]string) {
	for i := range lemin.Paths {
		var currComb [][]string
		currComb = append(currComb, lemin.Paths[i])

		for j, path := range lemin.Paths {
			if i == j {
				continue
			}
			isInter := false
			for _, currCombPath := range currComb {
				if pathsIntersect(currCombPath, path) {
					isInter = true
					break
				}
			}
			if !isInter {
				currComb = append(currComb, path)
			}
		}
		validCombs = append(validCombs, currComb)
	}
	return validCombs
}

func pathsIntersect(path1, path2 []string) bool {
	for _, roomName := range path1[1 : len(path1)-1] {
		if containsRoom(path2, roomName) {
			return true
		}
	}
	return false
}

func containsRoom(path []string, roomName string) bool {
	for _, room := range path[1 : len(path)-1] {
		if room == roomName {
			return true
		}
	}
	return false
}
