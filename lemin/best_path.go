package lemin

import (
	"fmt"
	"os"

	"Lemin-Project/variables"
)

// Determines the optimal combination of paths
func BestPath(ants int, combined [][][]string) [][]string {
	for _, path := range Paths {
		if len(path) == 2 {
			for i := 1; i <= ants; i++ {
				fmt.Print("L", i, "-", variables.End, " ")
			}
			fmt.Println("")
			os.Exit(0)
		}
	}

	var min int
	heights := make(map[int][][]string)
	for _, cmb := range combined {
		paths := len(cmb)
		max := len(cmb[len(cmb)-1])
		fill := 0
		for _, path := range cmb {
			fill += max - len(path)
		}
		left := ants - fill
		if left == 0 {
			continue
		}
		height := max + left/paths
		if left%paths != 0 {
			height++
		}

		heights[height] = cmb
		min = height
	}

	for h := range heights {
		if h < min {
			min = h
		}
	}

	return heights[min]
}
