package main

import (
	"fmt"
	"lemin/lemin"
)

func main() {
	//parse input, validate, find paths, simulate movements
	Variables := lemin.ReadFile()
	fmt.Println(Variables)
}
