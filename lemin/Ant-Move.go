package lemin

import (
	"fmt"

	"Lemin-Project/variables"
)

// Distributes ants to the respective paths
func AntMove(combinations [][]string, num int) {
	levels := num / len(combinations)
	if num%len(combinations) != 0 {
		levels++
	}

	ants := make([]variables.Ant, num+1)
	ants[0].IsInactive = true
	id := 0

	for j := 0; j < levels; j++ {
		for _, path := range combinations {
			id++
			ants[id].Path = path
			ants[id].CurrentRoom = 1
			ants[id].IsInactive = false
			if id == num {
				break
			}
		}
	}

	exit := false
	taken := make(map[string]bool)
	for !exit {
		for id, ant := range ants {
			if ant.IsInactive {
				continue
			}
			room := ant.Path[ant.CurrentRoom]
			if taken[room] {
				fmt.Println()
				break
			}
			fmt.Print("L", id, "-", room, " ")
			if id == num {
				fmt.Println()
				if room == variables.End {
					exit = true
				}
			}
			ants[id].CurrentRoom++
			taken[ants[id].Previous] = false
			if room != variables.End {
				taken[room] = true
				ants[id].Previous = room
			}
			if room == variables.End {
				ants[id].IsInactive = true
			}
		}
	}
}
