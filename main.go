package main

import (
	"Lemin-Project/lemin"
	"Lemin-Project/variables"
	"fmt"
)

func PrintInfos() {
	fmt.Print("Start room : " + variables.Start)
	fmt.Print("\n-----------------------------")
	fmt.Print("\nEnd room : " + variables.End)
	fmt.Println("\n-----------------------------")
	for key, value := range variables.Rooms {
		fmt.Printf("name: %v,  adrr: %p,  x: %v, y: %v,visited: %v, Relations: [", key, value, value.X, value.Y, value.Visited)
		for _, rel := range value.Relations {
			for key, value := range variables.Rooms {
				if value == rel {
					fmt.Printf(" %v - %p\t", key, value)
				}
			}
		}
		fmt.Println("]\n-----------------------------")
	}

	for i, v := range variables.Paths {
		fmt.Println("\n-----------------------------")
		fmt.Printf("Path %d : ", i)
		for _, c := range v {
			for n, r := range variables.Rooms {
				if c == r {
					fmt.Print(n, " ")
				}
			}
			fmt.Print()
		}
		fmt.Println("\n-----------------------------")
	}
}
func main() {
	lemin.FindAll()
	PrintInfos()
}
