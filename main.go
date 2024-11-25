package main

import (
	"Lemin-Project/lemin"
	"Lemin-Project/variables"
	"fmt"
)
func PrintInfos() {
	for key, value := range variables.Rooms {
		fmt.Printf("name: %v,  adrr: %p,  x: %v, y: %v,visited: %v, Relations: [", key, value, value.X, value.Y, value.Visited )
		for _, rel := range value.Relations {
			for key, value := range variables.Rooms {
				if value == rel {
					fmt.Printf("%v - %p\t", key, value)
				}
			}
		}
		fmt.Println("]\n-----------------------------")
	}
}
func main() {
	lemin.FindAll(&variables.Rooms)
	fmt.Println(variables.Paths)
}
