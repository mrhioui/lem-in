package main

import (
	"fmt"

	"Lemin-Project/lemin"
	"Lemin-Project/variables"
)

func main() {
	// Parse all info from file
	input := lemin.ParsingAll()

	// Extract and sort paths
	lemin.FindPaths(variables.Rooms[variables.Start], nil)
	lemin.SortPaths(variables.Paths)

	// Group paths in combinations
	combined := lemin.ValidPathsCombs()
	lemin.SortInCombinations(combined)

	fmt.Print(input, "\n\n")

	// Process with optimal combination
	combination := lemin.BestComb(variables.NAnt, combined)
	lemin.AntMove(combination, variables.NAnt)
}
