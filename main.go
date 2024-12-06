package main

import (
	"fmt"

	"Lemin-Project/lemin"
	"Lemin-Project/variables"
)

func main() {
	// Parse all info from file
	input := lemin.ParsingAll(&variables.Rooms)

	// Extract and sort paths
	lemin.FindPaths(variables.Rooms[variables.Start], nil)
	lemin.SortPaths(lemin.Paths)

	// Group paths in combinations
	combined := lemin.ValidPathsCombs()
	lemin.SortInCombinations(combined)

	fmt.Print(input, "\n\n")

	// Process with optimal combination
	combinations := lemin.BestPath(variables.NAnt, combined)
	lemin.AntMove(combinations, variables.NAnt)
}
