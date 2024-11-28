package main

import (
	"fmt"

	"Lemin-Project/lemin"
	"Lemin-Project/variables"
)

func PrintInfos() {
	for key, value := range variables.Rooms {
		fmt.Printf("name: %v,  adrr: %p,  x: %v, y: %v, visited: %v\n\tRelations: [", key, value, value.X, value.Y, value.Visited)
		for _, rel := range value.Relations {
			for key, value := range variables.Rooms {
				if value == rel {
					fmt.Printf("%v - %p\t", key, value)
				}
			}
		}
		fmt.Println("]\n-------------------------------------------------------------------------------------------------------------------------------------------------")
	}
}

func main() {
	lemin.FindAll(&variables.Rooms)
	PrintInfos()
	fmt.Println("\n-------------------------------------------------------------------------------------------------------------------------------------------------")
	lemin.ExtractPaths(variables.Rooms[variables.Start], nil)
	// lemin.ExtractPaths(variables.Rooms[variables.Start],[]*variables.Room{})
	lemin.SortPaths()
	for i, path := range lemin.Paths {
		fmt.Printf("path n%d: %v >> capcity: %d\n", i+1, path, len(path)-2)
	}
	fmt.Println("\n-------------------------------------------------------------------------------------------------------------------------------------------------")
	combined := ValidPathsCombs()
	for i, com := range combined {
		fmt.Printf("combination n%d: %v >> number of combined paths %d\n", i+1, com, len(com))
	}
	fmt.Println("\n-------------------------------------------------------------------------------------------------------------------------------------------------")
	// for i, path := range lemin.Paths {
	// 	fmt.Printf("path n%d: %v >> capcity: %d\n", i+1, path, len(path)-2)
	// }
}

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
