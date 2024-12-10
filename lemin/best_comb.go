package lemin

import (
	"fmt"
	"os"

	"Lemin-Project/variables"
)

// Determines optimal combination of paths based on turns
func BestComb(ants int, combined [][][]string) [][]string {
	// Start and End rooms are connected
	for _, path := range variables.Paths {
		if len(path) == 2 {
			for i := 1; i <= ants; i++ {
				fmt.Print("L", i, "-", variables.End, " ")
			}
			fmt.Println("")
			os.Exit(0)
		}
	}

	var optimal [][]string
	minTurns := int(^uint(0) >> 1) // Initialize to max int value

	for _, cmb := range combined {
		paths := len(cmb)           // Number of paths
		max := len(cmb[len(cmb)-1]) // Length of the longest path
		fill := 0                   // Ants needed to equalize paths

		// Calculate the fill required to equalize all paths to max
		for _, path := range cmb {
			fill += max - len(path)
		}

		// Remaining ants to distribute
		left := ants - fill
		if left < 0 {
			continue // Skip invalid combinations
		}

		// Calculate total number of turns
		turns := max + left/paths // Floored division: left/paths

		// Update the optimal combination if less turns are found
		if turns < minTurns {
			minTurns = turns
			optimal = cmb
		}
	}

	return optimal
}
