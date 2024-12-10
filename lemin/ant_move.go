package lemin

import (
	"fmt"
	"strings"
)

// Distributes ants to the respective paths
func AntMove(cmb [][]string, num int) {
	positions := make([]int, num)
	occupied := make(map[string]bool)
	startPaths := len(cmb)

	// Initializes positions
	for i := 0; i < num; i++ {
		positions[i] = 1
	}

	for {
		allAntsFinished := true
		roundOutput := []string{}
		antsMoving := 0

		for i := 0; i < num; i++ {
			if processAnt(i, cmb, positions, occupied,
				startPaths, &antsMoving, &roundOutput) {
				allAntsFinished = false
			}
		}

		if len(roundOutput) > 0 {
			fmt.Println(strings.Join(roundOutput, " "))
		}

		if allAntsFinished {
			break
		}
	}
}

// Processes an individual ant and updates its position
func processAnt(
	i int,
	cmb [][]string,
	positions []int,
	occupied map[string]bool,
	startPaths int,
	antsMoving *int,
	roundOutput *[]string,
) bool {
	pathId := i % len(cmb)
	if i == len(positions)-1 {
		pathId = 0 // Last ant follows the first path
	}

	if positions[i] >= len(cmb[pathId]) {
		return false // Ant has completed its path
	}

	next := cmb[pathId][positions[i]]

	// Free the previous position
	if positions[i] > 0 && positions[i]-1 < len(cmb[pathId]) {
		occupied[cmb[pathId][positions[i]-1]] = false
	}

	// Check if the ant can move from the start
	if positions[i] == 1 && *antsMoving >= startPaths {
		return true
	}

	// Move the ant if the next position is not occupied or is the endpoint
	if !occupied[next] || next == cmb[pathId][len(cmb[pathId])-1] {
		*roundOutput = append(*roundOutput, fmt.Sprintf("L%d-%s", i+1, next))
		occupied[next] = true
		positions[i]++
		if positions[i] == 2 {
			(*antsMoving)++
		}
	}

	return positions[i] < len(cmb[pathId])
}
